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
		// 'unevaluatedProperties: false' next to a '$ref' is the one permitted use:
		// it is the only way to reject unknown keys on a referenced object in draft
		// 2020-12. See isClosedRef.
		if s.UnevaluatedProperties != nil && !isClosedRef(s) {
			ruleResults.Add(
				"Property must not use unevaluatedProperties, except as 'unevaluatedProperties: false' next to '$ref'",
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
