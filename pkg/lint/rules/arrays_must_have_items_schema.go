package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type ArraysMustHaveItems struct{}

func (r ArraysMustHaveItems) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	callback := func(s *schema.ExtendedSchema) {
		if !hasItems(s) {
			ruleResults.Add(
				"Array must specify the schema of its items",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseArrays(s, callback)

	return *ruleResults
}

func hasItems(s *schema.ExtendedSchema) bool {
	return s.Items2020 != nil || s.Items != nil
}

func (r ArraysMustHaveItems) GetSeverity() Severity {
	return SeverityError
}
