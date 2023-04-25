package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type AvoidRecursion struct{}

func (r AvoidRecursion) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	recurse.RecurseAllPre(s, func(s *schema.ExtendedSchema) {
		if s.IsSelfReference() {
			ruleResults.Add(
				"Schema must not reference itself",
				s.GetResolvedLocation(),
			)
		}
	})
	return *ruleResults
}

func (r AvoidRecursion) GetSeverity() Severity {
	return SeverityError
}
