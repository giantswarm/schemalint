package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
	"github.com/giantswarm/schemalint/v2/pkg/lint/utils"
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

type DeprecatedPropertiesShouldHaveComment struct{}

func (r DeprecatedPropertiesShouldHaveComment) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	callback := func(schema *schemautils.ExtendedSchema) {
		if schema.Deprecated && schema.Comment == "" {
			ruleResults.Add(fmt.Sprintf("Deprecated property '%s' should have a $comment.", schema.GetHumanReadableLocation()), schema.GetResolvedLocation())
		}
	}

	utils.RecurseProperties(schema, callback)

	return *ruleResults
}

func (r DeprecatedPropertiesShouldHaveComment) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
