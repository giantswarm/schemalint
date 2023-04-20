package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type AvoidUnevaluated struct{}

func (r AvoidUnevaluated) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	callback := func(s *schema.ExtendedSchema) {
		if s.UnevaluatedItems != nil {
			ruleResults.Add(
				"Property must not use unevaluatedItems",
				s.GetResolvedLocation(),
			)
		}
		if s.UnevaluatedProperties != nil {
			ruleResults.Add(
				"Property must not use unevaluatedProperties",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseAll(s, callback)

	return *ruleResults
}

func (r AvoidUnevaluated) GetSeverity() Severity {
	return SeverityError
}
