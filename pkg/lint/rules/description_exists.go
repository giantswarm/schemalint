package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionExists struct{}

func (r DescriptionExists) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	callback := func(schema *schemautils.ExtendedSchema) {
		checkDescriptionExists(schema, ruleResults)
	}
	utils.RecurseProperties(schema, callback)
	return *ruleResults
}

func checkDescriptionExists(schema *schemautils.ExtendedSchema, ruleResults *lint.RuleResults) {
	if schema.Description == "" {
		ruleResults.Add(fmt.Sprintf("Property '%s' should have a description.", schema.GetConciseLocation()))
	}
}

func (r DescriptionExists) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
