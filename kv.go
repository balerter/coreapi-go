package coreapi

import (
	"encoding/json"
)

type ModuleKV struct {
	rf requestFunc
}

// Put the value to the storage
func (kv ModuleKV) Put(key, value string) error {
	_, err := kv.rf("kv/put/"+key, "text/plain", []byte(value))
	return err
}

// Upsert the value in the storage
func (kv ModuleKV) Upsert(key, value string) error {
	_, err := kv.rf("kv/upsert/"+key, "text/plain", []byte(value))
	return err
}

// Delete the value from the storage
func (kv ModuleKV) Delete(key string) error {
	_, err := kv.rf("kv/delete/"+key, "", nil)
	return err
}

// Get a value from the storage
func (kv ModuleKV) Get(key string) (string, error) {
	resp, err := kv.rf("kv/get/"+key, "", nil)
	if err != nil {
		return "", err
	}
	var result string
	errUnmarshal := json.Unmarshal(resp, &result)
	if errUnmarshal != nil {
		return "", errUnmarshal
	}
	return result, nil
}

// All returns all the values from the storage
func (kv ModuleKV) All() (map[string]string, error) {
	resp, err := kv.rf("kv/all", "", nil)
	if err != nil {
		return nil, err
	}
	result := map[string]string{}
	errUnmarshal := json.Unmarshal(resp, &result)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}
	return result, nil
}
