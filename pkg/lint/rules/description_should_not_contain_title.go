package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionShouldNotContainTitle struct{}

func (r DescriptionShouldNotContainTitle) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		checkDescriptionShouldNotContainTitle(schema, ruleResults)
	}

	utils.RecursePropertiesWithDescription(schema, callback)

	return *ruleResults
}

func checkDescriptionShouldNotContainTitle(schema *schemautils.ExtendedSchema, ruleResults *lint.RuleResults) {
	if schema.Title == "" {
		return
	}

	if strings.Contains(strings.ToLower(schema.Description), strings.ToLower(schema.Title)) {
		ruleResults.Add(fmt.Sprintf("Property '%s' description should not repeat the title.", schema.GetConciseLocation()))
	}
}

func (r DescriptionShouldNotContainTitle) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
