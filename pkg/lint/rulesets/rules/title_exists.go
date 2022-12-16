package rules

import (
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint/findings"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

// Check recursively that all properties have a title
type TitleExists struct{}

func (r TitleExists) Verify(schema *jsonschema.Schema) []RuleViolation {
	violations := []RuleViolation{}
	for _, property := range schema.Properties {
		violations = append(violations, VerifyR(property)...)
	}

	return violations

}

func VerifyR(schema *jsonschema.Schema) []RuleViolation {
	violations := []RuleViolation{}
	if schema.Title == "" {
		violations = append(violations, RuleViolation{
			Reason: "title is missing @ " + trimLocation(schema.Location),
		})
	}

	for _, property := range schema.Properties {
		violations = append(violations, VerifyR(property)...)
	}

	return violations
}

func (r TitleExists) GetSeverity() findings.Severity {
	return findings.SeverityError
}

func trimLocation(location string) string {
	return strings.Split(location, "#")[1]
}
