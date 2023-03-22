package utils

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type StringWithLevel struct {
	Value          string
	ReferenceLevel int
}

type InterfaceWithLevel struct {
	Value          []interface{}
	ReferenceLevel int
}

// Annotations for a property with the reference level at which each was found.
// The reference level describes how many '$ref's were resolved at the given
// property path.
// 0 is the root schema.
type AnnotationsWithLevel struct {
	Title       *StringWithLevel
	Description *StringWithLevel
	Examples    *InterfaceWithLevel
}

func (a *AnnotationsWithLevel) GetTitle() string {
	if a == nil || a.Title == nil {
		return ""
	}
	return a.Title.Value
}

func (a *AnnotationsWithLevel) GetDescription() string {
	if a == nil || a.Description == nil {
		return ""
	}
	return a.Description.Value
}

func (a *AnnotationsWithLevel) GetExamples() []interface{} {
	if a == nil || a.Examples == nil {
		return nil
	}
	return a.Examples.Value
}

func (a *AnnotationsWithLevel) UpdateAnnotationsIfNecessary(schema *schemautils.ExtendedSchema, level int) {
	if schema.Title != "" && (a.Title == nil || level <= a.Title.ReferenceLevel) {
		a.Title = &StringWithLevel{Value: schema.Title, ReferenceLevel: level}
	}
	if schema.Description != "" && (a.Description == nil || level <= a.Description.ReferenceLevel) {
		a.Description = &StringWithLevel{Value: schema.Description, ReferenceLevel: level}
	}
	if schema.Examples != nil && (a.Examples == nil || level <= a.Examples.ReferenceLevel) {
		a.Examples = &InterfaceWithLevel{Value: schema.Examples, ReferenceLevel: level}
	}
}

type PropertyAnnotationsMap map[string]*AnnotationsWithLevel

func NewPropertyAnnotationsMap() PropertyAnnotationsMap {
	return make(PropertyAnnotationsMap)
}

func (pam PropertyAnnotationsMap) UpdateAnnotationsIfNecessary(schema *schemautils.ExtendedSchema, level int) {
	location := schema.GetResolvedLocation()
	annotations, ok := pam[location]
	if !ok {
		annotations = &AnnotationsWithLevel{Title: nil, Description: nil, Examples: nil}
		pam[location] = annotations
	}
	annotations.UpdateAnnotationsIfNecessary(schema, level)
}

func (pam PropertyAnnotationsMap) WhereDescriptionsExist() PropertyAnnotationsMap {
	newMap := NewPropertyAnnotationsMap()
	for path, annotations := range pam {
		if annotations.GetDescription() != "" {
			newMap[path] = annotations
		}
	}
	return newMap
}

func (pam PropertyAnnotationsMap) WhereTitlesExist() PropertyAnnotationsMap {
	newMap := NewPropertyAnnotationsMap()
	for path, annotations := range pam {
		if annotations.GetTitle() != "" {
			newMap[path] = annotations
		}
	}
	return newMap
}

func (pam PropertyAnnotationsMap) GetParentAnnotations(resolvedLocation string) (*AnnotationsWithLevel, error) {
	parentResolvedLocation, err := schemautils.GetParentPropertyPath(resolvedLocation)

	if err != nil {
		return nil, err
	}

	annotations, ok := pam[parentResolvedLocation]
	if !ok {
		return nil, fmt.Errorf("Could not find parent annotations for %s", resolvedLocation)
	}

	return annotations, nil
}

// Builds a map with all properties in the given schema, where the key is the
// path to the property and the value are the annotations for that property.
// <path> -> <annotations>
func BuildPropertyAnnotationsMap(schema *schemautils.ExtendedSchema) PropertyAnnotationsMap {
	propertyAnnotationsMap := make(PropertyAnnotationsMap)
	RecurseProperties(schema, func(schema *schemautils.ExtendedSchema) {
		if schema.IsProperty() {
			referenceLevel := schema.GetReferenceLevel()
			propertyAnnotationsMap.UpdateAnnotationsIfNecessary(schema, referenceLevel)
		}

	})
	return propertyAnnotationsMap
}
