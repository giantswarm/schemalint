// Package normalize providesa a function to process a JSON input and return it in
// normalized form.
package normalize

import (
	"bytes"
	"encoding/json"
	"reflect"
)

const (
	// Four spaces indentation.
	indentation = "    "
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
	output, err := marshalIndentWithoutEscape(data, prefix, indentation)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func marshalIndentWithoutEscape(t interface{}, prefix, indent string) ([]byte, error) {
	marshalled, err := marshalWithoutEscape(t)

	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	err = json.Indent(&buffer, marshalled, prefix, indent)

	return buffer.Bytes(), err
}

func marshalWithoutEscape(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)

	return buffer.Bytes(), err
}

func CheckIsNormalized(jsonBytes []byte) (bool, error) {
	normalized, err := Normalize(jsonBytes)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(jsonBytes, normalized), nil
}
