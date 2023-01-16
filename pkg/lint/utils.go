package lint

import (
	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/schemautils"
)

func RecurseAll(schema *jsonschema.Schema, getViolationsAtLocation func(schema *jsonschema.Schema) []string) []string {
	violations := getViolationsAtLocation(schema)

	for _, property := range schema.Properties {
		violations = append(violations, RecurseAll(property, getViolationsAtLocation)...)
	}

	if schema.Items2020 != nil {
		violations = append(violations, RecurseAll(schema.Items2020, getViolationsAtLocation)...)
	}

	return violations
}

func RecurseProperties(schema *jsonschema.Schema, getViolationsAtLocation func(schema *jsonschema.Schema) []string) []string {
	getViolationsAtLocationIfProperty := func(schema *jsonschema.Schema) []string {
		if schemautils.IsProperty(schema) {
			return getViolationsAtLocation(schema)
		}
		return []string{}
	}

	return RecurseAll(schema, getViolationsAtLocationIfProperty)
}

func RecursePropertiesWithDescription(schema *jsonschema.Schema, getViolationsAtLocation func(schema *jsonschema.Schema) []string) []string {
	getViolationsAtLocationIfPropertyWithDescription := func(schema *jsonschema.Schema) []string {
		if schemautils.IsProperty(schema) && schema.Description != "" {
			return getViolationsAtLocation(schema)
		}
		return []string{}
	}

	return RecurseAll(schema, getViolationsAtLocationIfPropertyWithDescription)
}
