package schemautils

import (
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type ExtendedSchema struct {
	*jsonschema.Schema
	Parent           *ExtendedSchema
	ParentPath       string
	TrimFromLocation string
	RootFilePath     string
}

func NewExtendedSchema(schema *jsonschema.Schema) *ExtendedSchema {
	return &ExtendedSchema{
		Schema: schema,
	}
}

func (schema *ExtendedSchema) GetRefSchema() *ExtendedSchema {
	if schema.Ref == nil {
		return nil
	}
	refSchema := NewExtendedSchema(schema.Ref)

	refSchema.Parent = schema
	refSchema.ParentPath = schema.GetResolvedLocation()
	refSchema.TrimFromLocation = removeIdFromLocation(schema.Ref.Location)

	return refSchema
}

func (schema *ExtendedSchema) GetProperties() []*ExtendedSchema {
	properties := []*ExtendedSchema{}
	for _, property := range schema.Properties {
		newSchema := NewExtendedSchema(property)
		newSchema.InheritParentFrom(schema)
		properties = append(properties, newSchema)
	}
	return properties
}

func (schema *ExtendedSchema) InheritParentFrom(other *ExtendedSchema) {
	schema.Parent = other.Parent
	schema.ParentPath = other.ParentPath
	schema.TrimFromLocation = other.TrimFromLocation
	schema.RootFilePath = other.RootFilePath
}

func (schema *ExtendedSchema) GetItems2020() *ExtendedSchema {
	if schema.Items2020 == nil {
		return nil
	}
	newSchema := NewExtendedSchema(schema.Items2020)
	newSchema.InheritParentFrom(schema)
	return newSchema
}

func (schema *ExtendedSchema) GetItems() []*ExtendedSchema {
	if schema.Items == nil {
		return nil
	}

	schemas := []*jsonschema.Schema{}
	switch schema.Items.(type) {
	case *jsonschema.Schema:
		schemas = append(schemas, schema.Items.(*jsonschema.Schema))
	case []*jsonschema.Schema:
		schemas = append(schemas, schema.Items.([]*jsonschema.Schema)...)
	}

	extendedSchemas := make([]*ExtendedSchema, len(schemas))
	for i, item := range schemas {
		newSchema := NewExtendedSchema(item)
		newSchema.InheritParentFrom(schema)
		extendedSchemas[i] = newSchema
	}

	return extendedSchemas
}

// Gets current location without '/properties' in path
// e.g. /properties/cluster/properties/azure/properties/credentialSecret/properties/namespace
// becomes /cluster/azure/credentialSecret/namespace
func (schema *ExtendedSchema) GetConciseLocation() string {
	location := schema.GetResolvedLocation()
	location = strings.ReplaceAll(location, "/properties", "")
	return location
}

func (schema *ExtendedSchema) GetHumanReadableLocation() string {
	location := schema.GetConciseLocation()
	if location == "" {
		location = "/"
	}
	return location
}

// Gets the location, excluding ids, including potential parent schemas
func (schema *ExtendedSchema) GetResolvedLocation() string {
	location := removeIdFromLocation(schema.Location)
	if schema.TrimFromLocation != "" {
		location = strings.Replace(location, schema.TrimFromLocation, "", 1)
	}
	return schema.ParentPath + location
}

func removeIdFromLocation(location string) string {
	return strings.Split(location, "#")[1]
}

func (schema *ExtendedSchema) IsProperty() bool {
	location := schema.GetResolvedLocation()
	path := strings.Split(location, "/")
	return len(path) > 1 && path[len(path)-2] == "properties"
}

func (schema *ExtendedSchema) IsObject() bool {
	isObject := false
	for _, t := range schema.Types {
		if t == "object" {
			isObject = true
		}
	}
	return isObject
}

func (schema *ExtendedSchema) IsSelfReference() bool {
	parentLocations := schema.getParentEntryLocations()
	for _, parentLocation := range parentLocations {
		if parentLocation == schema.Location {
			return true
		}
	}
	return false
}

func (schema *ExtendedSchema) getParentEntryLocations() []string {
	locations := []string{}
	if schema.Parent != nil {
		locations = append(locations, schema.Parent.Location)
		locations = append(locations, schema.Parent.getParentEntryLocations()...)
	}
	return locations
}

func (schema *ExtendedSchema) GetReferenceLevel() int {
	if schema.Parent == nil {
		return 0
	}
	return schema.Parent.GetReferenceLevel() + 1
}

func GetParentPropertyPath(conciseLocation string) string {
	path := strings.Split(conciseLocation, "/")
	if len(path) <= 1 {
		return ""
	}
	return strings.Join(path[:len(path)-1], "/")
}
