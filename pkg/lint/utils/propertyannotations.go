package utils

import (
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
	if schema.Title != "" && (a.Title == nil || level < a.Title.ReferenceLevel) {
		a.Title = &StringWithLevel{Value: schema.Title, ReferenceLevel: level}
	}
	if schema.Description != "" && (a.Description == nil || level < a.Description.ReferenceLevel) {
		a.Description = &StringWithLevel{Value: schema.Description, ReferenceLevel: level}
	}
	if schema.Examples != nil && (a.Examples == nil || level < a.Examples.ReferenceLevel) {
		a.Examples = &InterfaceWithLevel{Value: schema.Examples, ReferenceLevel: level}
	}
}

type PropertyAnnotationsMap map[string]*AnnotationsWithLevel

func NewPropertyAnnotationsMap() PropertyAnnotationsMap {
	return make(PropertyAnnotationsMap)
}

func (pam PropertyAnnotationsMap) UpdateAnnotationsIfNecessary(schema *schemautils.ExtendedSchema, level int) {
	path := schema.GetConciseLocation()
	annotations, ok := pam[path]
	if !ok {
		annotations = &AnnotationsWithLevel{Title: nil, Description: nil, Examples: nil}
		pam[path] = annotations
	}
	annotations.UpdateAnnotationsIfNecessary(schema, level)
}

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
