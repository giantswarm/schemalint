package utils

import (
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type TitleWithLevel struct {
	Title          string
	ReferenceLevel int
}

type DescriptionWithLevel struct {
	Description    string
	ReferenceLevel int
}

type ExamplesWithLevel struct {
	Examples       []interface{}
	ReferenceLevel int
}

type AnnotationsWithLevel struct {
	Title       *TitleWithLevel
	Description *DescriptionWithLevel
	Examples    *ExamplesWithLevel
}

func (a *AnnotationsWithLevel) UpdateAnnotationsIfNecessary(schema *schemautils.ExtendedSchema, level int) {
	if a.Title == nil || level < a.Title.ReferenceLevel {
		a.Title = &TitleWithLevel{Title: schema.Title, ReferenceLevel: level}
	}
	if a.Description == nil || level < a.Description.ReferenceLevel {
		a.Description = &DescriptionWithLevel{Description: schema.Description, ReferenceLevel: level}
	}
	if a.Examples == nil || level < a.Examples.ReferenceLevel {
		a.Examples = &ExamplesWithLevel{Examples: schema.Examples, ReferenceLevel: level}
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
		annotations = &AnnotationsWithLevel{}
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
