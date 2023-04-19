package rules

import (
	"github.com/giantswarm/schemalint/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/pkg/schema"
)

type AvoidLogicalConstruct struct{}

func (r AvoidLogicalConstruct) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := RuleResults{}

	callback := func(s *schema.ExtendedSchema) {
		if s.If != nil || s.Then != nil || s.Else != nil {
			ruleResults.Add(
				"Schema must not use logical constructs (if, then, else)",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseAll(s, callback)

	return ruleResults
}

func (r AvoidLogicalConstruct) GetSeverity() Severity {
	return SeverityError
}
