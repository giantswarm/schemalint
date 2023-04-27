package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type AvoidRecursionKeywords struct{}

func (r AvoidRecursionKeywords) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	recurse.RecurseAll(s, func(s *schema.ExtendedSchema) {
		if s.DynamicAnchor != "" || s.DynamicRef != nil || s.RecursiveRef != nil {
			ruleResults.Add(
				"Schema must not use recursion keywords (dynamicAnchor, dynamicRef, recursiveRef)",
				s.GetResolvedLocation(),
			)
		}
	})
	return *ruleResults
}

func (r AvoidRecursionKeywords) GetSeverity() Severity {
	return SeverityError
}
