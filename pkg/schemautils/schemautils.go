package schemautils

import (
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

// Gets current location without '/properties' in path
func GetConciseLocation(schema *jsonschema.Schema) string {
	location := GetLocation(schema)
	// remove all occurrences of '/properties'
	location = strings.ReplaceAll(location, "/properties", "")
	return location
}

func GetLocation(schema *jsonschema.Schema) string {
	return TrimLocation(schema.Location)
}

func TrimLocation(location string) string {
	return strings.Split(location, "#")[1]
}

func IsProperty(schema *jsonschema.Schema) bool {
	location := GetLocation(schema)
	path := strings.Split(location, "/")
	return len(path) > 1 && path[len(path)-2] == "properties"
}
