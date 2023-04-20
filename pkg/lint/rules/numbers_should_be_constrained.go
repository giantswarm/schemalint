package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type NumbersShouldBeConstrained struct{}

func (r NumbersShouldBeConstrained) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	callback := func(s *schema.ExtendedSchema) {
		if s.Minimum == nil &&
			s.Maximum == nil &&
			s.ExclusiveMinimum == nil &&
			s.ExclusiveMaximum == nil {
			ruleResults.Add(
				"Numeric property should be constrained through 'minimum', 'maximum', 'exclusiveMinimum' or 'exclusiveMaximum'",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseNumerics(s, callback)
	return *ruleResults
}

func (r NumbersShouldBeConstrained) GetSeverity() Severity {
	return SeverityRecommendation
}
