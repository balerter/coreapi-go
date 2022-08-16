package coreapi

import (
	"fmt"
	"testing"
)

func TestModuleLog_Error(t *testing.T) {
	m := ModuleLog{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "log/error" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "message" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}

		return nil, fmt.Errorf("err1")
	}}

	err := m.Error("message")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleLog_Warn(t *testing.T) {
	m := ModuleLog{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "log/warn" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "message" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}

		return nil, fmt.Errorf("err1")
	}}

	err := m.Warn("message")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleLog_Info(t *testing.T) {
	m := ModuleLog{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "log/info" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "message" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}

		return nil, fmt.Errorf("err1")
	}}

	err := m.Info("message")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleLog_Debug(t *testing.T) {
	m := ModuleLog{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "log/debug" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "message" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}

		return nil, fmt.Errorf("err1")
	}}

	err := m.Debug("message")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}
