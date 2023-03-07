package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type StringsShouldBeConstrained struct{}

func (r StringsShouldBeConstrained) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	callback := func(schema *schemautils.ExtendedSchema) {
		if schema.Pattern == nil &&
			schema.MinLength == -1 &&
			schema.MaxLength == -1 &&
			schema.Enum == nil &&
			schema.Constant == nil &&
			schema.Format == "" {
			ruleResults.Add(
				fmt.Sprintf(
					"String property '%s' should be constrained through 'pattern', 'minLength', 'maxLength', 'enum', 'constant' or 'format'.",
					schema.GetHumanReadableLocation()),
				schema.GetResolvedLocation(),
			)
		}
	}

	utils.RecurseStrings(schema, callback)

	return *ruleResults
}

func (r StringsShouldBeConstrained) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
