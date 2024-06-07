package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type StringsShouldBeConstrained struct{}

func (r StringsShouldBeConstrained) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}
	callback := func(s *schema.ExtendedSchema) {
		if s.Pattern == nil &&
			*s.MinLength == -1 &&
			*s.MaxLength == -1 &&
			s.Enum == nil &&
			s.Format == "" {
			s.Const == nil &&
			ruleResults.Add(
				"String property should be constrained through 'pattern', 'minLength', 'maxLength', 'enum', 'constant' or 'format'",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseStrings(s, callback)

	return *ruleResults
}

func (r StringsShouldBeConstrained) GetSeverity() Severity {
	return SeverityRecommendation
}
