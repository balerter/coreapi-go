package coreapi

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestModuleChart_Render_error_call_rf(t *testing.T) {
	m := ModuleChart{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "chart/render" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "application/json" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != `{"title":"t","series":[{"Color":"1","LineColor":"2","PointColor":"3","Data":[{"Timestamp":10,"Value":20}]}]}` {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return nil, fmt.Errorf("err1")
		},
	}

	_, err := m.Render("t", []DataSeries{{
		Color:      "1",
		LineColor:  "2",
		PointColor: "3",
		Data: []DataItem{{
			Timestamp: 10,
			Value:     20,
		}},
	}})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to call chart/render: err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleChart_Render_error_unmarshal_response(t *testing.T) {
	m := ModuleChart{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "chart/render" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "application/json" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != `{"title":"t","series":[{"Color":"1","LineColor":"2","PointColor":"3","Data":[{"Timestamp":10,"Value":20}]}]}` {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte("bad"), nil
		},
	}

	_, err := m.Render("t", []DataSeries{{
		Color:      "1",
		LineColor:  "2",
		PointColor: "3",
		Data: []DataItem{{
			Timestamp: 10,
			Value:     20,
		}},
	}})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "unmarshal response error, invalid character 'b' looking for beginning of value" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleChart_Render_error_base64_decode(t *testing.T) {
	m := ModuleChart{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "chart/render" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "application/json" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != `{"title":"t","series":[{"Color":"1","LineColor":"2","PointColor":"3","Data":[{"Timestamp":10,"Value":20}]}]}` {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte(`"aaa"`), nil
		},
	}

	_, err := m.Render("t", []DataSeries{{
		Color:      "1",
		LineColor:  "2",
		PointColor: "3",
		Data: []DataItem{{
			Timestamp: 10,
			Value:     20,
		}},
	}})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "decode response error, illegal base64 data at input byte 0" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleChart_Render(t *testing.T) {
	m := ModuleChart{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "chart/render" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "application/json" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != `{"title":"t","series":[{"Color":"1","LineColor":"2","PointColor":"3","Data":[{"Timestamp":10,"Value":20}]}]}` {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte(`"` + base64.RawStdEncoding.EncodeToString([]byte("xxx")) + `"`), nil
		},
	}

	resp, err := m.Render("t", []DataSeries{{
		Color:      "1",
		LineColor:  "2",
		PointColor: "3",
		Data: []DataItem{{
			Timestamp: 10,
			Value:     20,
		}},
	}})
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if string(resp) != "xxx" {
		t.Fatalf("unexpected response, got %s", resp)
	}
}
