package lint

import "github.com/santhosh-tekuri/jsonschema/v5"

func RecurseCall(schema *jsonschema.Schema, getViolationsAtLocation func(schema *jsonschema.Schema) []string) []string {
	violations := getViolationsAtLocation(schema)

	for _, property := range schema.Properties {
		violations = append(violations, RecurseCall(property, getViolationsAtLocation)...)
	}

	if schema.Items2020 != nil {
		violations = append(violations, RecurseCall(schema.Items2020, getViolationsAtLocation)...)
	}

	return violations
}
