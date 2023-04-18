package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
	"github.com/giantswarm/schemalint/v2/pkg/lint/utils"
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

type NumbersShouldBeConstrained struct{}

func (r NumbersShouldBeConstrained) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		if schema.Minimum == nil &&
			schema.Maximum == nil &&
			schema.ExclusiveMinimum == nil &&
			schema.ExclusiveMaximum == nil {
			ruleResults.Add(
				fmt.Sprintf(
					"Numeric property '%s' should be constrained through 'minimum', 'maximum', 'exclusiveMinimum' or 'exclusiveMaximum'",
					schema.GetHumanReadableLocation(),
				), schema.GetResolvedLocation(),
			)
		}
	}

	utils.RecurseNumerics(schema, callback)
	return *ruleResults
}

func (r NumbersShouldBeConstrained) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
