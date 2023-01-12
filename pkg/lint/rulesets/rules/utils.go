package rules

import "github.com/santhosh-tekuri/jsonschema/v5"

func recurseCall(schema *jsonschema.Schema, getViolationsAtLocation func(schema *jsonschema.Schema) []RuleViolation) []RuleViolation {
	violations := getViolationsAtLocation(schema)

	for _, property := range schema.Properties {
		violations = append(violations, recurseCall(property, getViolationsAtLocation)...)
	}

	if schema.Items2020 != nil {
		violations = append(violations, recurseCall(schema.Items2020, getViolationsAtLocation)...)
	}

	return violations
}
