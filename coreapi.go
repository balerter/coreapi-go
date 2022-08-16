package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type apiResponse struct {
	Status string          `json:"status"`
	Error  string          `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Balerter is the main struct for the coreapi package.
type Balerter struct {
	// Alert provides access to the alert module
	Alert ModuleAlert
	// Datasource provides access to the datasource module
	Datasource ModuleDatasource
	// KV provides access to the kv module
	KV ModuleKV
	// Log provides access to the log module
	Log ModuleLog
	// TLS provides access to the tls module
	TLS ModuleTLS
	// Runtime provides access to the runtime module
	Runtime ModuleRuntime
	// Chart provides access to the chart module
	Chart ModuleChart

	address   string
	authToken string
	client    httpClient
}

type requestFunc func(path, contentType string, body []byte) ([]byte, error)

// New creates a new Balerter instance.
// address is the address of the balerter server.
// authToken is the authentication token for the balerter server. Pass empty token, if auth is not use.
func New(address, authToken string) *Balerter {
	c := &Balerter{
		authToken: authToken,
		address:   strings.Trim(address, "/"),
		client:    &http.Client{},
	}
	c.Alert = ModuleAlert{rf: c.request}
	c.Datasource = ModuleDatasource{rf: c.request}
	c.KV = ModuleKV{rf: c.request}
	c.Log = ModuleLog{rf: c.request}
	c.TLS = ModuleTLS{rf: c.request}
	c.Runtime = ModuleRuntime{rf: c.request}
	c.Chart = ModuleChart{rf: c.request}

	return c
}

func (b *Balerter) request(path, contentType string, body []byte) ([]byte, error) {
	u := fmt.Sprintf("%s/%s", b.address, path)

	req, errReq := http.NewRequest(http.MethodPost, u, bytes.NewReader(body))
	if errReq != nil {
		return nil, errReq
	}
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}
	if b.authToken != "" {
		req.Header.Add("Authorization", b.authToken)
	}

	resp, errDo := b.client.Do(req)
	if errDo != nil {
		return nil, errDo
	}

	defer resp.Body.Close()

	r := apiResponse{}

	errRead := json.NewDecoder(resp.Body).Decode(&r)
	if errRead != nil {
		return nil, fmt.Errorf("error decode response, %w", errRead)
	}

	if r.Status == "error" {
		return nil, fmt.Errorf("%s", r.Error)
	}

	return r.Result, nil
}
