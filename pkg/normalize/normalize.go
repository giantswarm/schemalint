// Package normalize providesa a function to process a JSON input and return it in
// normalized form.
package normalize

import (
	"encoding/json"
	"reflect"
)

const (
	// Four spaces indentation.
	indendation = "    "
	prefix      = ""
)

// Normalize takes JSON and returns normalized JSON.
func Normalize(jsonBytes []byte) ([]byte, error) {
	var data interface{}
	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, err
	}

	// We use the fact that MarshalIndent applies a specific sorting
	// of object keys, apart from normalizing indentation.
	output, err := json.MarshalIndent(data, prefix, indendation)
	if err != nil {
		return nil, err
	}

	// trailing newline
	output = append(output, []byte("\n")...)

	return output, nil
}

func CheckIsNormalized(jsonBytes []byte) (bool, error) {
	normalized, err := Normalize(jsonBytes)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(jsonBytes, normalized), nil
}
