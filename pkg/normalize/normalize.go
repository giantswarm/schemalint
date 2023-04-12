// Package normalize providesa a function to process a JSON input and return it in
// normalized form.
package normalize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/iancoleman/orderedmap"
)

const arrayWrapperKey = "arrayWrapperKey"

// Normalize takes JSON and returns normalized JSON.
//
// Normalization includes:
//   - Sorting keys according to the configured importance.
//   - Consistent indentation.
func Normalize(jsonBytes []byte) ([]byte, error) {
	data, err := loadToOrderedMap(jsonBytes)
	if err != nil {
		return nil, err
	}

	lessFunc := func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		aImportance := getKeyImportance(a.Key(), a.Value())
		bImportance := getKeyImportance(b.Key(), b.Value())

		return aImportance > bImportance
	}

	deepSortOrderedMap(data, lessFunc)

	output, err := marshalIndentWithoutEscape(data, prefix, indentation)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func loadToOrderedMap(jsonBytes []byte) (interface{}, error) {
	data := initOrderedMap()

	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		iData, err := loadArrayToOrderedMap(jsonBytes)
		if err != nil {
			return nil, err
		}
		return iData, nil
	}

	return data, nil
}

// loadArrayToOrderedMap is a workaround for the fact that the orderedmap library
// does not support unmarshalling into an array.
//
// This functions adds a dummy key around the array and then removes it again.
func loadArrayToOrderedMap(jsonBytes []byte) (interface{}, error) {
	wrappedData := initOrderedMap()

	arrayBytes := []byte(fmt.Sprintf(`{"%s":%s}`, arrayWrapperKey, jsonBytes))

	err := json.Unmarshal(arrayBytes, &wrappedData)
	if err != nil {
		return nil, err
	}

	data, ok := wrappedData.Get(arrayWrapperKey)

	if !ok {
		return nil, fmt.Errorf("arrayWrapperKey not found")
	}

	return data, nil
}

func initOrderedMap() *orderedmap.OrderedMap {
	data := orderedmap.New()
	data.SetEscapeHTML(false)
	return data
}

func getKeyImportance(key string, value interface{}) int {
	usedKeyImportance := getKeyImportanceMap()

	for _, override := range getKeyImportanceMapOverrides() {
		if override.When(key, value) {
			usedKeyImportance = mergeMaps(usedKeyImportance, override.KeyImportance)
		}
	}

	importance, ok := usedKeyImportance[key]
	if !ok {
		return defaultKeyImportance
	}
	return importance
}

func mergeMaps(maps ...map[string]int) map[string]int {
	result := make(map[string]int)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func deepSortOrderedMap(
	data interface{},
	lessFunc func(a *orderedmap.Pair, b *orderedmap.Pair) bool,
) {

	switch data.(type) {
	case *orderedmap.OrderedMap:
		o, _ := data.(*orderedmap.OrderedMap)
		o.Sort(lessFunc)
		for _, k := range o.Keys() {
			v, _ := o.Get(k)

			if vMap, ok := v.(orderedmap.OrderedMap); ok {
				deepSortOrderedMap(&vMap, lessFunc)
			}
		}
	case orderedmap.OrderedMap:
		o, _ := data.(orderedmap.OrderedMap)
		o.Sort(lessFunc)
		for _, k := range o.Keys() {
			v, _ := o.Get(k)

			if vMap, ok := v.(orderedmap.OrderedMap); ok {
				deepSortOrderedMap(&vMap, lessFunc)
			}
		}
	case []interface{}:
		for _, v := range data.([]interface{}) {
			if vMap, ok := v.(orderedmap.OrderedMap); ok {
				deepSortOrderedMap(&vMap, lessFunc)
			}
		}
	default:
		panic(fmt.Sprintf("unexpected type: %T", data))
	}

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
