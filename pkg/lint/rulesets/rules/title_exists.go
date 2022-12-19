package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint/findings"
	"github.com/giantswarm/schemalint/pkg/schemautils"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

// Check recursively that all properties have a title
type TitleExists struct{}

func (r TitleExists) Verify(schema *jsonschema.Schema) []RuleViolation {
	return recurseCall(schema, checkTitle)
}

func checkTitle(schema *jsonschema.Schema) []RuleViolation {
	ruleViolations := []RuleViolation{}

	if !schemautils.IsProperty(schema) {
		return ruleViolations
	}

	if schema.Title == "" {
		ruleViolations = append(ruleViolations, RuleViolation{
			Reason: fmt.Sprintf("Property '%s' has no title", schemautils.GetLocation(schema)),
		})
	}
	return ruleViolations
}

func (r TitleExists) GetSeverity() findings.Severity {
	return findings.SeverityError
}
