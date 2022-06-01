package loader

import (
	"encoding/json"
	"io/ioutil"

	// Package to strip UTF-8 BOM
	"bytes"
)

// This package should Read and parse a JSON file and strip UTF-8 BOM
func LoadJsonFile(path string) (map[string]interface{}, error) {
	// Variable used to store the JSON object on json.Unmarshal
	var jsonObject map[string]interface{}

	// Read the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Strip UTF-8 BOM
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	// Parse the JSON
	err = json.Unmarshal(data, &jsonObject)
	if err != nil {
		return nil, err
	}

	return jsonObject, nil
}
