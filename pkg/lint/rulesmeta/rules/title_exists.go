package rules

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint/rulesmeta"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

// Check recursively that all properties have a title
type TitleExists struct{}

func (r TitleExists) Verify(schema *jsonschema.Schema) []string {
	return rulesmeta.RecurseCall(schema, checkTitle)
}

func checkTitle(schema *jsonschema.Schema) []string {
	ruleViolations := []string{}

	if !schemautils.IsProperty(schema) {
		return ruleViolations
	}

	if schema.Title == "" {
		ruleViolations = append(ruleViolations, fmt.Sprintf("Property '%s' must have a title.", schemautils.GetConciseLocation(schema)))
	}
	return ruleViolations
}

func (r TitleExists) GetSeverity() rulesmeta.Severity {
	return rulesmeta.SeverityError
}
