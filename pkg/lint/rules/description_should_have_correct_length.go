package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

const maxDescriptionLength = 200
const minDescriptionLength = 50

type DescriptionShouldHaveCorrectLength struct{}

func (r DescriptionShouldHaveCorrectLength) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		checkDescriptionShouldHaveCorrectLength(schema, ruleResults)
	}

	utils.RecursePropertiesWithDescription(schema, callback)
	return *ruleResults
}

func checkDescriptionShouldHaveCorrectLength(schema *schemautils.ExtendedSchema, ruleResults *lint.RuleResults) {
	if len(schema.Description) > maxDescriptionLength {
		ruleResults.Add(fmt.Sprintf("Property '%s' description should be less than %d characters.", schema.GetConciseLocation(), maxDescriptionLength))
	}

	if len(schema.Description) < minDescriptionLength {
		ruleResults.Add(fmt.Sprintf("Property '%s' description should be more than %d characters.", schema.GetConciseLocation(), minDescriptionLength))
	}
}

func (r DescriptionShouldHaveCorrectLength) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
