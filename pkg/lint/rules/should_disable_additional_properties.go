package rules

import (
	"github.com/giantswarm/schemalint/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/pkg/schema"
)

type ShouldDisableAdditionalProperties struct{}

func (r ShouldDisableAdditionalProperties) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}

	callback := func(s *schema.ExtendedSchema) {
		if !isAdditionalPropertiesDisabled(s) {
			ruleResults.Add(
				"Object should disable additional properties",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseObjects(s, callback)

	return *ruleResults
}

func isAdditionalPropertiesDisabled(s *schema.ExtendedSchema) bool {
	return s.AdditionalProperties == false
}

func (r ShouldDisableAdditionalProperties) GetSeverity() Severity {
	return SeverityRecommendation
}
