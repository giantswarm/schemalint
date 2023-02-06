package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DeprecatedPropertiesShouldHaveComment struct{}

func (r DeprecatedPropertiesShouldHaveComment) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	callback := func(schema *schemautils.ExtendedSchema) {
		if schema.Deprecated && schema.Comment == "" {
			ruleResults.Add(fmt.Sprintf("Deprecated property '%s' should have a $comment.", schema.GetHumanReadableLocation()))
		}
	}

	utils.RecurseProperties(schema, callback)

	return *ruleResults
}

func (r DeprecatedPropertiesShouldHaveComment) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
