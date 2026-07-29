package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
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
	// An object defined via '$ref' is closed with 'unevaluatedProperties: false'
	// instead -- 'additionalProperties: false' would reject every field the
	// referenced schema defines. See isClosedRef. The '$ref' target is exempt too:
	// closing it would reject the sibling properties the referring schema adds.
	return s.AdditionalProperties == false || isClosedRef(s) || isClosedRefTarget(s)
}

func (r ShouldDisableAdditionalProperties) GetSeverity() Severity {
	return SeverityRecommendation
}
