package schemautils

import (
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

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
