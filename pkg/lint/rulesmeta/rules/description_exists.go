package rules

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint/rulesmeta"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionExists struct{}

func (r DescriptionExists) Verify(schema *jsonschema.Schema) []string {
	return rulesmeta.RecurseCall(schema, checkDescription)
}

func checkDescription(schema *jsonschema.Schema) []string {
	ruleViolations := []string{}

	if !schemautils.IsProperty(schema) {
		return ruleViolations
	}

	if schema.Description == "" {
		ruleViolations = append(ruleViolations, fmt.Sprintf("Property '%s' should have a description.", schemautils.GetConciseLocation(schema)))
	}
	return ruleViolations
}

func (r DescriptionExists) GetSeverity() rulesmeta.Severity {
	return rulesmeta.SeverityRecomendation
}