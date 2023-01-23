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
	refSchema.TrimFromLocation = RemoveIdFromLocation(schema.Ref.Location)

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
}

func (schema *ExtendedSchema) GetItems2020() *ExtendedSchema {
	if schema.Items2020 == nil {
		return nil
	}
	newSchema := NewExtendedSchema(schema.Items2020)
	newSchema.InheritParentFrom(schema)
	return newSchema
}

// Gets current location without '/properties' in path
// e.g. /properties/cluster/properties/azure/properties/credentialSecret/properties/namespace
// becomes /cluster/azure/credentialSecret/namespace
func (schema *ExtendedSchema) GetConciseLocation() string {
	location := schema.GetResolvedLocation()
	location = strings.ReplaceAll(location, "/properties", "")
	return location
}

// Gets the location, excluding ids, including potential parent schemas
func (schema *ExtendedSchema) GetResolvedLocation() string {
	location := RemoveIdFromLocation(schema.Location)
	if schema.TrimFromLocation != "" {
		location = strings.Replace(location, schema.TrimFromLocation, "", 1)
	}
	return schema.ParentPath + location
}

func RemoveIdFromLocation(location string) string {
	return strings.Split(location, "#")[1]
}

func (schema *ExtendedSchema) IsProperty() bool {
	location := schema.GetResolvedLocation()
	path := strings.Split(location, "/")
	return len(path) > 1 && path[len(path)-2] == "properties"
}

func (schema *ExtendedSchema) IsSelfReference() bool {
	parentLocations := schema.getParentLocations()
	for _, parentLocation := range parentLocations {
		if parentLocation == schema.Location {
			return true
		}
	}
	return false
}

func (schema *ExtendedSchema) getParentLocations() []string {
	locations := []string{}
	if schema.Parent != nil {
		locations = append(locations, schema.Parent.Location)
		locations = append(locations, schema.Parent.getParentLocations()...)
	}
	return locations
}

func (schema *ExtendedSchema) GetReferenceLevel() int {
	if schema.Parent == nil {
		return 0
	}
	return schema.Parent.GetReferenceLevel() + 1
}
