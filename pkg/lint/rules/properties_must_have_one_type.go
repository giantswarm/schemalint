package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type PropertiesMustHaveOneType struct{}

func (r PropertiesMustHaveOneType) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	callback := func(s *schema.ExtendedSchema) {
		numTypes := len(s.Types.ToStrings())
		if numTypes > 1 || (numTypes == 0 && s.Ref == nil) {
			ruleResults.Add(
				"Property must have exactly one type or be a reference to another schema and not have a type",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseProperties(s, callback)
	return *ruleResults
}

func (r PropertiesMustHaveOneType) GetSeverity() Severity {
	return SeverityError
}
