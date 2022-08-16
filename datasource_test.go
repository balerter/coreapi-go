package coreapi

import (
	"fmt"
	"testing"
)

func TestModuleDatasource_MySQL_error_call_rf(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/mysql/test/query" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return nil, fmt.Errorf("err1")
		},
	}
	mysql := m.MySQL("test")
	if mysql.name != "test" {
		t.Fatalf("expected name to be test, got %s", mysql.name)
	}
	_, err := mysql.Query("query")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to call mysql.query: err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleDatasource_MySQL(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/mysql/test/query" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte("foo"), nil
		},
	}
	mysql := m.MySQL("test")
	if mysql.name != "test" {
		t.Fatalf("expected name to be test, got %s", mysql.name)
	}
	resp, err := mysql.Query("query")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if string(resp) != "foo" {
		t.Fatalf("unexpected response, got %s", string(resp))
	}
}

func TestModuleDatasource_Loki_query_error_call_rf(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/loki/test/query" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return nil, fmt.Errorf("err1")
		},
	}
	loki := m.Loki("test")
	if loki.name != "test" {
		t.Fatalf("expected name to be test, got %s", loki.name)
	}
	_, err := loki.Query("query", nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to call loki.query: err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleDatasource_Loki_query(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/loki/test/query?direction=bar&limit=10&time=20" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte("foo"), nil
		},
	}
	loki := m.Loki("test")
	if loki.name != "test" {
		t.Fatalf("expected name to be test, got %s", loki.name)
	}
	resp, err := loki.Query("query", &LokiQueryParams{
		Limit:     10,
		Time:      20,
		Direction: "bar",
	})
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if string(resp) != `foo` {
		t.Fatalf("unexpected response, got %s", string(resp))
	}
}

func TestModuleDatasource_Loki_range_error_call_rf(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/loki/test/range" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return nil, fmt.Errorf("err1")
		},
	}
	loki := m.Loki("test")
	if loki.name != "test" {
		t.Fatalf("expected name to be test, got %s", loki.name)
	}
	_, err := loki.Range("query", nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to call loki.range: err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleDatasource_Loki_range(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/loki/test/range?direction=bar&end=30&limit=10&start=20&step=40" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte("foo"), nil
		},
	}
	loki := m.Loki("test")
	if loki.name != "test" {
		t.Fatalf("expected name to be test, got %s", loki.name)
	}
	resp, err := loki.Range("query", &LokiRangeParams{
		Limit:     10,
		Start:     20,
		End:       30,
		Step:      40,
		Direction: "bar",
	})
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if string(resp) != "foo" {
		t.Fatalf("unexpected response, got %s", string(resp))
	}
}

func TestModuleDatasource_Postgres_error_call_rf(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/postgres/test/query" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return nil, fmt.Errorf("err1")
		},
	}
	p := m.Postgres("test")
	if p.name != "test" {
		t.Fatalf("expected name to be test, got %s", p.name)
	}
	_, err := p.Query("query")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to call postgres.query: err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleDatasource_Postgres(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/postgres/test/query" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte("foo"), nil
		},
	}
	p := m.Postgres("test")
	if p.name != "test" {
		t.Fatalf("expected name to be test, got %s", p.name)
	}
	resp, err := p.Query("query")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if string(resp) != "foo" {
		t.Fatalf("unexpected response, got %s", string(resp))
	}
}

func TestModuleDatasource_Clickhouse_error_call_rf(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/clickhouse/test/query" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return nil, fmt.Errorf("err1")
		},
	}
	p := m.Clickhouse("test")
	if p.name != "test" {
		t.Fatalf("expected name to be test, got %s", p.name)
	}
	_, err := p.Query("query")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "failed to call clickhouse.query: err1" {
		t.Fatalf("unexpected error value, got %s", err.Error())
	}
}

func TestModuleDatasource_Clickhouse(t *testing.T) {
	m := ModuleDatasource{
		rf: func(path, contentType string, body []byte) ([]byte, error) {
			if path != "datasource/clickhouse/test/query" {
				t.Fatalf("unexpected path value, got %s", path)
			}
			if contentType != "text/plain" {
				t.Fatalf("unexpected contentType value, got %s", contentType)
			}
			if string(body) != "query" {
				t.Fatalf("unexpected body value, got %s", string(body))
			}
			return []byte("foo"), nil
		},
	}
	p := m.Clickhouse("test")
	if p.name != "test" {
		t.Fatalf("expected name to be test, got %s", p.name)
	}
	resp, err := p.Query("query")
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if string(resp) != "foo" {
		t.Fatalf("unexpected response, got %s", string(resp))
	}
}
