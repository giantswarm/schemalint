package rules

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint/findings"
	"github.com/giantswarm/schemalint/pkg/lint/rulesmeta"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionExists struct{}

func (r DescriptionExists) Verify(schema *jsonschema.Schema) []rulesmeta.RuleViolation {
	return rulesmeta.RecurseCall(schema, checkDescription)
}

func checkDescription(schema *jsonschema.Schema) []rulesmeta.RuleViolation {
	ruleViolations := []rulesmeta.RuleViolation{}

	if !schemautils.IsProperty(schema) {
		return ruleViolations
	}

	if schema.Description == "" {
		ruleViolations = append(ruleViolations, rulesmeta.RuleViolation{
			Reason: fmt.Sprintf("Property '%s' has no description", schemautils.GetLocation(schema)),
		})
	}
	return ruleViolations
}

func (r DescriptionExists) GetSeverity() findings.Severity {
	return findings.SeverityInfo
}
