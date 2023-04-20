package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type DeprecatedPropertiesShouldHaveComment struct{}

func (r DeprecatedPropertiesShouldHaveComment) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}
	callback := func(s *schema.ExtendedSchema) {
		if s.Deprecated && s.Comment == "" {
			ruleResults.Add(
				"Deprecated property should have a $comment",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseProperties(s, callback)

	return *ruleResults
}

func (r DeprecatedPropertiesShouldHaveComment) GetSeverity() Severity {
	return SeverityRecommendation
}
