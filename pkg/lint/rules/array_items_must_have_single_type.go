package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type ArrayItemsMustHaveSingleType struct{}

func (r ArrayItemsMustHaveSingleType) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	callback := func(s *schema.ExtendedSchema) {
		if containsIllegalArrayKeywords(s) {
			ruleResults.Add(
				"Array must not use illegal keyword(s): 'additionalItems', 'contains', 'prefixItems'",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseArrays(s, callback)

	return *ruleResults
}

func containsIllegalArrayKeywords(s *schema.ExtendedSchema) bool {
	if s.AdditionalItems != nil || s.Contains != nil || s.PrefixItems != nil {
		return true
	}
	return false
}

func (r ArrayItemsMustHaveSingleType) GetSeverity() Severity {
	return SeverityError
}
