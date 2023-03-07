package schemautils

import (
	"fmt"
	"strings"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
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
	properties := make([]*ExtendedSchema, 0, len(schema.Properties))
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

func (schema *ExtendedSchema) GetAnyOf() []*ExtendedSchema {
	if schema.AnyOf == nil {
		return nil
	}
	newSchemas := make([]*ExtendedSchema, len(schema.AnyOf))
	for i, anyOf := range schema.AnyOf {
		newSchema := NewExtendedSchema(anyOf)
		newSchema.InheritParentFrom(schema)
		newSchemas[i] = newSchema
	}
	return newSchemas
}

func (schema *ExtendedSchema) GetOneOf() []*ExtendedSchema {
	if schema.OneOf == nil {
		return nil
	}
	newSchemas := make([]*ExtendedSchema, len(schema.OneOf))
	for i, oneOf := range schema.OneOf {
		newSchema := NewExtendedSchema(oneOf)
		newSchema.InheritParentFrom(schema)
		newSchemas[i] = newSchema
	}
	return newSchemas
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
	location = ConvertToConciseLocation(location)
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
	return locationIsProperty(location)
}

func locationIsProperty(resolvedLocation string) bool {
	path := strings.Split(resolvedLocation, "/")
	return len(path) > 1 && path[len(path)-2] == "properties"
}

func (schema *ExtendedSchema) IsObject() bool {
	return schema.isType("object")
}

func (schema *ExtendedSchema) IsArray() bool {
	return schema.isType("array")
}

func (schema *ExtendedSchema) IsString() bool {
	return schema.isType("string")
}

func (schema *ExtendedSchema) IsNumber() bool {
	return schema.isType("number")
}

func (schema *ExtendedSchema) IsInteger() bool {
	return schema.isType("integer")
}

func (schema *ExtendedSchema) IsNumeric() bool {
	return schema.IsNumber() || schema.IsInteger()
}

func (schema *ExtendedSchema) IsBoolean() bool {
	return schema.isType("boolean")
}

func (schema *ExtendedSchema) isType(typeName string) bool {
	isType := false
	for _, t := range schema.Types {
		if t == typeName {
			isType = true
		}
	}
	return isType
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

func GetParentPropertyPath(resolvedLocation string) (string, error) {
	if !locationIsProperty(resolvedLocation) {
		return "", fmt.Errorf("location is not a property: %s", resolvedLocation)
	}

	path := strings.Split(resolvedLocation, "/")
	if len(path) <= 2 {
		return "", nil
	}
	return strings.Join(path[:len(path)-2], "/"), nil
}

// Returns a list of all schemas, whose location matches the given location.
// Due to the usage of '$ref', multiple schemas can have the same path.
func (schema *ExtendedSchema) GetSchemasAt(resolvedLocation string) []*ExtendedSchema {
	schemas := []*ExtendedSchema{}
	currentResolvedLocation := schema.GetResolvedLocation()
	if currentResolvedLocation == resolvedLocation {
		schemas = append(schemas, schema)
	}
	for _, property := range schema.GetProperties() {
		schemas = append(schemas, property.GetSchemasAt(resolvedLocation)...)
	}
	for _, item := range schema.GetItems() {
		schemas = append(schemas, item.GetSchemasAt(resolvedLocation)...)
	}
	if schema.GetRefSchema() != nil {
		schemas = append(schemas, schema.GetRefSchema().GetSchemasAt(resolvedLocation)...)
	}
	return schemas
}

func IsChildLocation(parentLocation string, childLocation string) bool {
	parentPaths := strings.Split(parentLocation, "/")
	childPaths := strings.Split(childLocation, "/")
	if len(childPaths) <= len(parentPaths) {
		return false
	}
	for i, parentPath := range parentPaths {
		if parentPath != childPaths[i] {
			return false
		}
	}
	return true
}

func ConvertToConciseLocation(resolvedLocation string) string {
	return strings.ReplaceAll(resolvedLocation, "/properties", "")
}
