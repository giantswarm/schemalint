package utils

import (
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

func RecurseAll(schema *schemautils.ExtendedSchema, getViolationsAtLocation func(schema *schemautils.ExtendedSchema) []string) []string {
	if schema.IsSelfReference() {
		return []string{}
	}

	violations := getViolationsAtLocation(schema)
	if schema.Ref != nil {
		refSchema := schema.GetRefSchema()
		violations = append(violations, RecurseAll(refSchema, getViolationsAtLocation)...)
	}

	for _, property := range schema.GetProperties() {
		violations = append(violations, RecurseAll(property, getViolationsAtLocation)...)
	}

	if schema.Items2020 != nil {
		violations = append(violations, RecurseAll(schema.GetItems2020(), getViolationsAtLocation)...)
	}

	return violations
}

func RecurseProperties(schema *schemautils.ExtendedSchema, getViolationsAtLocation func(schema *schemautils.ExtendedSchema) []string) []string {
	getViolationsAtLocationIfProperty := func(schema *schemautils.ExtendedSchema) []string {
		if schema.IsProperty() {
			return getViolationsAtLocation(schema)
		}
		return []string{}
	}

	return RecurseAll(schema, getViolationsAtLocationIfProperty)
}

func RecursePropertiesWithDescription(schema *schemautils.ExtendedSchema, getViolationsAtLocation func(schema *schemautils.ExtendedSchema) []string) []string {
	getViolationsAtLocationIfPropertyWithDescription := func(schema *schemautils.ExtendedSchema) []string {
		if schema.IsProperty() && schema.Description != "" {
			return getViolationsAtLocation(schema)
		}
		return []string{}
	}

	return RecurseAll(schema, getViolationsAtLocationIfPropertyWithDescription)
}
