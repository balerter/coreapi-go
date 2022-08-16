package coreapi

import (
	"encoding/json"
	"fmt"
)

type ModuleTLS struct {
	rf requestFunc
}

type TLSResult struct {
	Issuer         string   `json:"issuer"`
	Expiry         int64    `json:"expiry"`
	DNSNames       []string `json:"dns_names"`
	EmailAddresses []string `json:"email_addresses"`
}

// Get returns TLS info for the hostname. You should to define a hostname without the protocol and/or the port
func (t *ModuleTLS) Get(hostname string) ([]TLSResult, error) {
	respBody, err := t.rf("tls/get", "text/plain", []byte(hostname))
	if err != nil {
		return nil, fmt.Errorf("failed to get tls info: %w", err)
	}

	var resp []TLSResult

	errUnmarshal := json.Unmarshal(respBody, &resp)
	if errUnmarshal != nil {
		return nil, fmt.Errorf("failed to unmarshal tls info: %w", errUnmarshal)
	}

	return resp, nil
}
