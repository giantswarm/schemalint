package schema

import (
	"fmt"
	"strings"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

type ExtendedSchema struct {
	*jsonschema.Schema
	Parent           *ExtendedSchema
	ParentPath       string
	TrimFromLocation string
	RootFilePath     string
}

func NewExtendedSchema(s *jsonschema.Schema) *ExtendedSchema {
	return &ExtendedSchema{
		Schema: s,
	}
}

func (s *ExtendedSchema) GetRefSchema() *ExtendedSchema {
	if s.Ref == nil {
		return nil
	}
	refSchema := NewExtendedSchema(s.Ref)

	refSchema.Parent = s
	refSchema.ParentPath = s.GetResolvedLocation()
	refSchema.TrimFromLocation = removeIdFromLocation(s.Ref.Location)

	return refSchema
}

func (s *ExtendedSchema) GetProperties() map[string]*ExtendedSchema {
	properties := map[string]*ExtendedSchema{}
	for key, property := range s.Properties {
		newSchema := NewExtendedSchema(property)
		newSchema.InheritParentFrom(s)
		properties[key] = newSchema
	}
	return properties
}

// GetAdditionalProperties returns nil or bool or *ExtendedSchema
func (s *ExtendedSchema) GetAdditionalProperties() interface{} {
	if s.AdditionalProperties == nil {
		return nil
	}
	if s.AdditionalProperties == true {
		return true
	}
	if s.AdditionalProperties == false {
		return false
	}
	return NewExtendedSchema(s.AdditionalProperties.(*jsonschema.Schema))
}

func (s *ExtendedSchema) InheritParentFrom(other *ExtendedSchema) {
	s.Parent = other.Parent
	s.ParentPath = other.ParentPath
	s.TrimFromLocation = other.TrimFromLocation
	s.RootFilePath = other.RootFilePath
}

func (s *ExtendedSchema) GetItems2020() *ExtendedSchema {
	if s.Items2020 == nil {
		return nil
	}
	newSchema := NewExtendedSchema(s.Items2020)
	newSchema.InheritParentFrom(s)
	return newSchema
}

func (s *ExtendedSchema) GetAnyOf() []*ExtendedSchema {
	if s.AnyOf == nil {
		return nil
	}
	newSchemas := make([]*ExtendedSchema, len(s.AnyOf))
	for i, anyOf := range s.AnyOf {
		newSchema := NewExtendedSchema(anyOf)
		newSchema.InheritParentFrom(s)
		newSchemas[i] = newSchema
	}
	return newSchemas
}

func (s *ExtendedSchema) GetOneOf() []*ExtendedSchema {
	if s.OneOf == nil {
		return nil
	}
	newSchemas := make([]*ExtendedSchema, len(s.OneOf))
	for i, oneOf := range s.OneOf {
		newSchema := NewExtendedSchema(oneOf)
		newSchema.InheritParentFrom(s)
		newSchemas[i] = newSchema
	}
	return newSchemas
}

func (s *ExtendedSchema) GetItems() []*ExtendedSchema {
	if s.Items == nil {
		return nil
	}

	schemas := []*jsonschema.Schema{}
	switch s.Items.(type) {
	case *jsonschema.Schema:
		schemas = append(schemas, s.Items.(*jsonschema.Schema))
	case []*jsonschema.Schema:
		schemas = append(schemas, s.Items.([]*jsonschema.Schema)...)
	}

	extendedSchemas := make([]*ExtendedSchema, len(schemas))
	for i, item := range schemas {
		newSchema := NewExtendedSchema(item)
		newSchema.InheritParentFrom(s)
		extendedSchemas[i] = newSchema
	}

	return extendedSchemas
}

// Gets current location without '/properties' in path
// e.g. /properties/cluster/properties/azure/properties/credentialSecret/properties/namespace
// becomes /cluster/azure/credentialSecret/namespace
func (s *ExtendedSchema) GetConciseLocation() string {
	location := s.GetResolvedLocation()
	location = ConvertToConciseLocation(location)
	return location
}

// Gets the location, excluding ids, including potential parent schemas.
// The root location is '/'.
func (s *ExtendedSchema) GetResolvedLocation() string {
	location := removeIdFromLocation(s.Location)
	if s.TrimFromLocation != "" {
		location = strings.Replace(location, s.TrimFromLocation, "", 1)
	}
	location = s.ParentPath + location
	if location == "" {
		location = "/"
	}
	return location
}

func removeIdFromLocation(location string) string {
	return strings.Split(location, "#")[1]
}

func (s *ExtendedSchema) IsProperty() bool {
	location := s.GetResolvedLocation()
	return locationIsProperty(location)
}

func locationIsProperty(resolvedLocation string) bool {
	path := strings.Split(resolvedLocation, "/")
	return len(path) > 1 && path[len(path)-2] == "properties"
}

func (s *ExtendedSchema) IsObject() bool {
	return s.IsType("object")
}

func (s *ExtendedSchema) IsArray() bool {
	return s.IsType("array")
}

func (s *ExtendedSchema) IsString() bool {
	return s.IsType("string")
}

func (s *ExtendedSchema) IsNumber() bool {
	return s.IsType("number")
}

func (s *ExtendedSchema) IsInteger() bool {
	return s.IsType("integer")
}

func (s *ExtendedSchema) IsNumeric() bool {
	return s.IsNumber() || s.IsInteger()
}

func (s *ExtendedSchema) IsBoolean() bool {
	return s.IsType("boolean")
}

func (s *ExtendedSchema) IsType(typeName string) bool {
	isType := false
	for _, t := range s.Types.ToStrings() {
		if t == typeName {
			isType = true
		}
	}
	return isType
}

func (s *ExtendedSchema) IsSelfReference() bool {
	parentLocations := s.getParentEntryLocations()
	for _, parentLocation := range parentLocations {
		if parentLocation == s.Location {
			return true
		}
	}
	return false
}

func (s *ExtendedSchema) getParentEntryLocations() []string {
	locations := []string{}
	if s.Parent != nil {
		locations = append(locations, s.Parent.Location)
		locations = append(locations, s.Parent.getParentEntryLocations()...)
	}
	return locations
}

func (s *ExtendedSchema) GetReferenceLevel() int {
	if s.Parent == nil {
		return 0
	}
	return s.Parent.GetReferenceLevel() + 1
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
// Due to the usage of '$ref', there can be multiple schema definitions for
// the same location.
//
// For more information on this is look at the '"Overriding" Properties and
// Understanding `PropertyAnnotationsMap`' section in the README.
func (s *ExtendedSchema) GetSchemasAt(resolvedLocation string) []*ExtendedSchema {
	schemas := []*ExtendedSchema{}
	currentResolvedLocation := s.GetResolvedLocation()
	if currentResolvedLocation == resolvedLocation {
		schemas = append(schemas, s)
	}
	for _, property := range s.GetProperties() {
		schemas = append(schemas, property.GetSchemasAt(resolvedLocation)...)
	}
	for _, item := range s.GetItems() {
		schemas = append(schemas, item.GetSchemasAt(resolvedLocation)...)
	}
	if s.GetRefSchema() != nil {
		schemas = append(schemas, s.GetRefSchema().GetSchemasAt(resolvedLocation)...)
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
