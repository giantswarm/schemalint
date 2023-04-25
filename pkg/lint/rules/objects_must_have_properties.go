package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type ObjectsMustHaveProperties struct{}

func (r ObjectsMustHaveProperties) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	callback := func(s *schema.ExtendedSchema) {
		nProperties := len(s.Properties) + len(s.PatternProperties)
		_, ok := s.GetAdditionalProperties().(*schema.ExtendedSchema)
		if ok {
			nProperties++
		}

		if nProperties == 0 {
			ruleResults.Add(
				"Object must have at least one property",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseObjects(s, callback)
	return *ruleResults
}

func (r ObjectsMustHaveProperties) GetSeverity() Severity {
	return SeverityError
}
