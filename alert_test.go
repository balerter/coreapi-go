package coreapi

import (
	"fmt"
	"testing"
	"time"
)

func TestModuleAlert_error(t *testing.T) {
	m := ModuleAlert{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "alert/success/a" && path != "alert/warn/a" && path != "alert/error/a" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "b" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return nil, fmt.Errorf("err1")
		},
	}

	_, _, err1 := m.Success("a", "b", nil)
	if err1 == nil {
		t.Fatalf("expected error, got nil")
	}
	if err1.Error() != "failed to call alert/success/a: err1" {
		t.Fatalf("unexpected error value, got %s", err1.Error())
	}

	_, _, err2 := m.Warning("a", "b", nil)
	if err2 == nil {
		t.Fatalf("expected error, got nil")
	}
	if err2.Error() != "failed to call alert/warn/a: err1" {
		t.Fatalf("unexpected error value, got %s", err2.Error())
	}

	_, _, err3 := m.Error("a", "b", nil)
	if err3 == nil {
		t.Fatalf("expected error, got nil")
	}
	if err3.Error() != "failed to call alert/error/a: err1" {
		t.Fatalf("unexpected error value, got %s", err3.Error())
	}
}

func TestModuleAlert_call_unmarshal_error(t *testing.T) {
	m := ModuleAlert{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "alert/a/b" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "c" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte("bad"), nil
		},
	}

	_, _, err := m.call("a", "b", "c", nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to unmarshal response: invalid character 'b' looking for beginning of value" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleAlert_call(t *testing.T) {
	alertStart := time.Now()
	alertLastChange := time.Now()

	m := ModuleAlert{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "alert/a/b?channels=foo%2Cbar&escalate=10%3Afoo3%2Cbar3&fields=foo2%3Abar2&image=image.png&quiet=true&repeat=42" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "c" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte(fmt.Sprintf(`{"level_was_updated":true,"alert":{"level":3,"start":%q,"count":5,"name":"aa","last_change":%q}}`, alertStart.Format(time.RFC3339Nano), alertLastChange.Format(time.RFC3339Nano))), nil
		},
	}

	a, wasUpdated, err := m.call("a", "b", "c", &AlertOptions{
		Channels: []string{"foo", "bar"},
		Quiet:    true,
		Repeat:   42,
		Image:    "image.png",
		Fields:   map[string]string{"foo2": "bar2"},
		Escalate: map[int][]string{10: {"foo3", "bar3"}},
	})

	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}
	if !wasUpdated {
		t.Fatalf("expected wasUpdated value, got false")
	}
	if a.Name != "aa" {
		t.Fatalf("unexpected level name, got %s", a.Name)
	}
	if a.Level != 3 {
		t.Fatalf("unexpected level value, got %d", a.Level)
	}
	if !alertStart.Equal(a.Start) {
		t.Fatalf("unexpected start value, got %s, expect %s", a.Start, alertStart)
	}
	if !alertLastChange.Equal(a.LastChange) {
		t.Fatalf("unexpected lastChange value, got %s, expect %s", a.LastChange, alertLastChange)
	}
	if a.Count != 5 {
		t.Fatalf("unexpected count value, got %d, expect 5", a.Count)
	}
}

func TestModuleAlert_Get_unmarshal_error(t *testing.T) {
	m := ModuleAlert{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			return []byte("bad"), nil
		},
	}

	_, err := m.Get("a")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to unmarshal response: invalid character 'b' looking for beginning of value" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleAlert_Get_error_call_rf(t *testing.T) {
	m := ModuleAlert{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			return nil, fmt.Errorf("err1")
		},
	}

	_, err := m.Get("a")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to call alert/get/a: err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleAlert_Get(t *testing.T) {
	alertStart := time.Now()
	alertLastChange := time.Now()

	m := ModuleAlert{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "alert/get/a" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if len(body) != 0 {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte(fmt.Sprintf(`{"level":3,"start":%q,"count":5,"name":"aa","last_change":%q}`, alertStart.Format(time.RFC3339Nano), alertLastChange.Format(time.RFC3339Nano))), nil
		},
	}

	a, err := m.Get("a")

	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}
	if a.Name != "aa" {
		t.Fatalf("unexpected level name, got %s", a.Name)
	}
	if a.Level != 3 {
		t.Fatalf("unexpected level value, got %d", a.Level)
	}
	if !alertStart.Equal(a.Start) {
		t.Fatalf("unexpected start value, got %s, expect %s", a.Start, alertStart)
	}
	if !alertLastChange.Equal(a.LastChange) {
		t.Fatalf("unexpected lastChange value, got %s, expect %s", a.LastChange, alertLastChange)
	}
	if a.Count != 5 {
		t.Fatalf("unexpected count value, got %d, expect 5", a.Count)
	}
}
