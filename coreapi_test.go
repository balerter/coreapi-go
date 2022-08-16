package coreapi

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	a := New("/a//", "t")
	if a.address != "a" {
		t.Errorf("expected address to be 'a', got %s", a.address)
	}
	if a.authToken != "t" {
		t.Errorf("expected authToken to be 't', got %s", a.authToken)
	}
}

func TestCoreAPI_request_error_create_create_request(t *testing.T) {
	m := Balerter{}

	_, err := m.request("\n", "text", []byte("body"))
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "parse \"/\\n\": net/url: invalid control character in URL" {
		t.Errorf("unexpected error value, got %s", err.Error())
	}
}

type httpClientMock struct {
	do func(req *http.Request) (*http.Response, error)
}

func (c *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	return c.do(req)
}

func TestCoreAPI_request_error_create_do_request(t *testing.T) {
	cl := &httpClientMock{
		do: func(req *http.Request) (*http.Response, error) {
			return nil, fmt.Errorf("err1")
		},
	}

	m := Balerter{
		client: cl,
	}

	_, err := m.request("foo", "text", []byte("body"))
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Errorf("unexpected error value, got %s", err.Error())
	}
}

func TestCoreAPI_request_error_decode_response(t *testing.T) {
	cl := &httpClientMock{
		do: func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				Body: io.NopCloser(strings.NewReader(`bad`)),
			}
			return resp, nil
		},
	}

	m := Balerter{
		client: cl,
	}

	_, err := m.request("foo", "text", []byte("body"))
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "error decode response, invalid character 'b' looking for beginning of value" {
		t.Errorf("unexpected error value, got %s", err.Error())
	}
}

func TestCoreAPI_request_error_status(t *testing.T) {
	cl := &httpClientMock{
		do: func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				Body: io.NopCloser(strings.NewReader(`{"status":"error","error":"err1"}`)),
			}
			return resp, nil
		},
	}

	m := Balerter{
		client: cl,
	}

	_, err := m.request("foo", "text", []byte("body"))
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "err1" {
		t.Errorf("unexpected error value, got %s", err.Error())
	}
}

func TestCoreAPI_request(t *testing.T) {
	cl := &httpClientMock{
		do: func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("content-type") != "text" {
				t.Fatalf("expected content-type to be 'text', got %s", req.Header.Get("content-type"))
			}
			if req.Header.Get("authorization") != "t" {
				t.Fatalf("expected authorization to be 't', got %s", req.Header.Get("authorization"))
			}

			resp := &http.Response{
				Body: io.NopCloser(strings.NewReader(`{"status":"success","result":"foobar"}`)),
			}
			return resp, nil
		},
	}

	m := Balerter{
		client:    cl,
		authToken: "t",
	}

	resp, err := m.request("foo", "text", []byte("body"))
	if err != nil {
		t.Error("unexpected error, %w", err)
	}

	if string(resp) != `"foobar"` {
		t.Errorf("unexpected response, got %s", string(resp))
	}
}
