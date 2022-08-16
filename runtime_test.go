package coreapi

import (
	"fmt"
	"testing"
)

func TestModuleRuntime_Get_error(t *testing.T) {
	m := ModuleRuntime{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "runtime/get" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return nil, fmt.Errorf("err1")
	}}

	_, err := m.Get()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleRuntime_Get_error_unmarshal(t *testing.T) {
	m := ModuleRuntime{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "runtime/get" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return []byte("xxx"), nil
	}}

	_, err := m.Get()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "invalid character 'x' looking for beginning of value" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleRuntime_Get(t *testing.T) {
	m := ModuleRuntime{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "runtime/get" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return []byte(`{"log_level":"l","is_debug":true,"is_once":true,"with_script":"s","config_source":"c","safe_mode":true}`), nil
	}}

	resp, err := m.Get()
	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}
	if resp.LogLevel != "l" {
		t.Fatalf("unexpected LogLevel value, got %s", resp.LogLevel)
	}
	if !resp.IsDebug {
		t.Fatalf("unexpected IsDebug value, got %t", resp.IsDebug)
	}
	if !resp.IsOnce {
		t.Fatalf("unexpected IsOnce value, got %t", resp.IsOnce)
	}
	if resp.WithScript != "s" {
		t.Fatalf("unexpected WithScript value, got %s", resp.WithScript)
	}
	if resp.ConfigSource != "c" {
		t.Fatalf("unexpected ConfigSource value, got %s", resp.ConfigSource)
	}
	if !resp.SafeMode {
		t.Fatalf("unexpected SafeMode value, got %t", resp.SafeMode)
	}
}
