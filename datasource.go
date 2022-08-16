package coreapi

import (
	"fmt"
	"net/url"
	"strconv"
)

type ModuleDatasource struct {
	rf requestFunc
}

// MySQL

// MySQL provides access to mysql datasources.
func (m ModuleDatasource) MySQL(name string) ModuleDatasourceMySQL {
	return ModuleDatasourceMySQL{rf: m.rf, name: name}
}

type ModuleDatasourceMySQL struct {
	rf   requestFunc
	name string
}

// Query method for the mysql datasource.
func (m ModuleDatasourceMySQL) Query(query string) ([]byte, error) {
	resp, err := m.rf("datasource/mysql/"+m.name+"/query", "text/plain", []byte(query))
	if err != nil {
		return nil, fmt.Errorf("failed to call mysql.query: %w", err)
	}
	return resp, nil
}

// Loki

// Loki provides access to loki datasources.
func (m ModuleDatasource) Loki(name string) ModuleDatasourceLoki {
	return ModuleDatasourceLoki{rf: m.rf, name: name}
}

type ModuleDatasourceLoki struct {
	rf   requestFunc
	name string
}

type LokiQueryParams struct {
	Limit     int
	Time      int
	Direction string
}

func (p *LokiQueryParams) toQuery() string {
	vals := url.Values{}
	if p.Limit != 0 {
		vals.Add("limit", strconv.Itoa(p.Limit))
	}
	if p.Time != 0 {
		vals.Add("time", strconv.Itoa(p.Time))
	}
	if p.Direction != "" {
		vals.Add("direction", p.Direction)
	}
	return vals.Encode()
}

// Query method for the loki datasource.
func (m ModuleDatasourceLoki) Query(query string, params *LokiQueryParams) ([]byte, error) {
	u := "datasource/loki/" + m.name + "/query"
	if params != nil {
		u += "?" + params.toQuery()
	}
	resp, err := m.rf(u, "text/plain", []byte(query))
	if err != nil {
		return nil, fmt.Errorf("failed to call loki.query: %w", err)
	}
	return resp, nil
}

type LokiRangeParams struct {
	Limit     int
	Start     int
	End       int
	Step      int
	Direction string
}

func (p *LokiRangeParams) toQuery() string {
	vals := url.Values{}
	if p.Limit != 0 {
		vals.Add("limit", strconv.Itoa(p.Limit))
	}
	if p.Start != 0 {
		vals.Add("start", strconv.Itoa(p.Start))
	}
	if p.End != 0 {
		vals.Add("end", strconv.Itoa(p.End))
	}
	if p.Step != 0 {
		vals.Add("step", strconv.Itoa(p.Step))
	}
	if p.Direction != "" {
		vals.Add("direction", p.Direction)
	}
	return vals.Encode()
}

// Range method for the loki datasource.
func (m ModuleDatasourceLoki) Range(query string, params *LokiRangeParams) ([]byte, error) {
	u := "datasource/loki/" + m.name + "/range"
	if params != nil {
		u += "?" + params.toQuery()
	}
	resp, err := m.rf(u, "text/plain", []byte(query))
	if err != nil {
		return nil, fmt.Errorf("failed to call loki.range: %w", err)
	}
	return resp, nil
}

// Postgres

// Postgres provides access to postgres datasources.
func (m ModuleDatasource) Postgres(name string) ModuleDatasourcePostgres {
	return ModuleDatasourcePostgres{rf: m.rf, name: name}
}

type ModuleDatasourcePostgres struct {
	rf   requestFunc
	name string
}

// Query method for the postgres datasource.
func (m ModuleDatasourcePostgres) Query(query string) ([]byte, error) {
	resp, err := m.rf("datasource/postgres/"+m.name+"/query", "text/plain", []byte(query))
	if err != nil {
		return nil, fmt.Errorf("failed to call postgres.query: %w", err)
	}
	return resp, nil
}

// Clickhouse

// Clickhouse provides access to clickhouse datasources.
func (m ModuleDatasource) Clickhouse(name string) ModuleDatasourceClickhouse {
	return ModuleDatasourceClickhouse{rf: m.rf, name: name}
}

type ModuleDatasourceClickhouse struct {
	rf   requestFunc
	name string
}

// Query method for the clickhouse datasource.
func (m ModuleDatasourceClickhouse) Query(query string) ([]byte, error) {
	resp, err := m.rf("datasource/clickhouse/"+m.name+"/query", "text/plain", []byte(query))
	if err != nil {
		return nil, fmt.Errorf("failed to call clickhouse.query: %w", err)
	}
	return resp, nil
}
