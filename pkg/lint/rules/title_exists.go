package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

// Check recursively that all properties have a title
type TitleExists struct{}

func (r TitleExists) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	callback := func(schema *schemautils.ExtendedSchema) {
		checkTitle(schema, ruleResults)
	}
	utils.RecurseProperties(schema, callback)
	return *ruleResults
}

func checkTitle(schema *schemautils.ExtendedSchema, ruleResults *lint.RuleResults) {
	if schema.Title == "" {
		ruleResults.Add(fmt.Sprintf("Property '%s' must have a title.", schema.GetConciseLocation()))
	}
}

func (r TitleExists) GetSeverity() lint.Severity {
	return lint.SeverityError
}
