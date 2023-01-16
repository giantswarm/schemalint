package rules

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint/findings"
	"github.com/giantswarm/schemalint/pkg/lint/rulesmeta"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

// Check recursively that all properties have a title
type TitleExists struct{}

func (r TitleExists) Verify(schema *jsonschema.Schema) []rulesmeta.RuleViolation {
	return rulesmeta.RecurseCall(schema, checkTitle)
}

func checkTitle(schema *jsonschema.Schema) []rulesmeta.RuleViolation {
	ruleViolations := []rulesmeta.RuleViolation{}

	if !schemautils.IsProperty(schema) {
		return ruleViolations
	}

	if schema.Title == "" {
		ruleViolations = append(ruleViolations, rulesmeta.RuleViolation{
			Reason: fmt.Sprintf("Property '%s' has no title", schemautils.GetLocation(schema)),
		})
	}
	return ruleViolations
}

func (r TitleExists) GetSeverity() findings.Severity {
	return findings.SeverityError
}
