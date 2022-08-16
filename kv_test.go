package coreapi

import (
	"fmt"
	"testing"
)

func TestModuleKV_Put_error(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/put/k" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "v" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return nil, fmt.Errorf("err1")
	}}

	err := m.Put("k", "v")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleKV_Put(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/put/k" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "v" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return nil, nil
	}}

	err := m.Put("k", "v")
	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}
}

func TestModuleKV_Upsert_error(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/upsert/k" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "v" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return nil, fmt.Errorf("err1")
	}}

	err := m.Upsert("k", "v")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleKV_Upsert(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/upsert/k" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "v" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return nil, nil
	}}

	err := m.Upsert("k", "v")
	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}
}

func TestModuleKV_Delete_error(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/delete/k" {
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

	err := m.Delete("k")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleKV_Delete(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/delete/k" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return nil, nil
	}}

	err := m.Delete("k")
	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}
}

func TestModuleKV_Get_error(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/get/k" {
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

	_, err := m.Get("k")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleKV_Get_error_unmarshal(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/get/k" {
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

	_, err := m.Get("k")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "invalid character 'x' looking for beginning of value" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleKV_Get(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/get/k" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return []byte(`"v"`), nil
	}}

	resp, err := m.Get("k")
	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}
	if resp != "v" {
		t.Fatalf("unexpected response value, got %s", resp)
	}
}

func TestModuleKV_All_error(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/all" {
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

	_, err := m.All()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}
func TestModuleKV_All_error_unmarshal(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/all" {
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

	_, err := m.All()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "invalid character 'x' looking for beginning of value" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleKV_All(t *testing.T) {
	m := ModuleKV{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "kv/all" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return []byte(`{"k1":"v1"}`), nil
	}}

	resp, err := m.All()
	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}
	if len(resp) != 1 {
		t.Fatalf("unexpected response value, got %s", resp)
	}
	v, ok := resp["k1"]
	if !ok {
		t.Fatalf("unexpected response value, got %s", resp)
	}
	if v != "v1" {
		t.Fatalf("unexpected response value, got %s", resp)
	}
}
