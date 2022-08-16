package coreapi

import (
	"encoding/json"
)

type ModuleRuntime struct {
	rf requestFunc
}

type RuntimeInfo struct {
	LogLevel     string `json:"log_level"`
	IsDebug      bool   `json:"is_debug"`
	IsOnce       bool   `json:"is_once"`
	WithScript   string `json:"with_script"`
	ConfigSource string `json:"config_source"`
	SafeMode     bool   `json:"safe_mode"`
}

// Get returns runtime info.
func (b *ModuleRuntime) Get() (*RuntimeInfo, error) {
	resp, err := b.rf("runtime/get", "", nil)
	if err != nil {
		return nil, err
	}
	info := RuntimeInfo{}
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
