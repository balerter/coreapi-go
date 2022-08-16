package coreapi

import (
	"fmt"
	"testing"
)

func TestModuleTLS_Get_error(t *testing.T) {
	m := ModuleTLS{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "tls/get" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "domain.com" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return nil, fmt.Errorf("err1")
	}}

	_, err := m.Get("domain.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to get tls info: err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleTLS_Get_error_unmarshal(t *testing.T) {
	m := ModuleTLS{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "tls/get" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "domain.com" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return []byte("xxx"), nil
	}}

	_, err := m.Get("domain.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to unmarshal tls info: invalid character 'x' looking for beginning of value" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleTLS_Get(t *testing.T) {
	m := ModuleTLS{rf: func(path, contentType string, body []byte) ([]byte, error) {
		if path != "tls/get" {
			t.Fatalf("unexpected path value, got %s", path)
		}
		if contentType != "text/plain" {
			t.Fatalf("unexpected contentType value, got %s", contentType)
		}
		if string(body) != "domain.com" {
			t.Fatalf("unexpected body value, got %s", string(body))
		}
		return []byte(`[{"issuer":"i","expiry":1,"dns_names":["n"],"email_addresses":["a"]}]`), nil
	}}

	resp, err := m.Get("domain.com")
	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}

	if len(resp) != 1 {
		t.Fatalf("unexpected response length, got %d", len(resp))
	}
	if resp[0].Issuer != "i" {
		t.Fatalf("unexpected issuer value, got %s", resp[0].Issuer)
	}
	if resp[0].Expiry != 1 {
		t.Fatalf("unexpected expiry value, got %d", resp[0].Expiry)
	}
	if len(resp[0].DNSNames) != 1 || resp[0].DNSNames[0] != "n" {
		t.Fatalf("unexpected dns_names value, got %v", resp[0].DNSNames)
	}
	if len(resp[0].EmailAddresses) != 1 || resp[0].EmailAddresses[0] != "a" {
		t.Fatalf("unexpected email_addresses value, got %v", resp[0].EmailAddresses)
	}
}
