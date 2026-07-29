package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type ObjectsMustHaveProperties struct{}

func (r ObjectsMustHaveProperties) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	callback := func(s *schema.ExtendedSchema) {
		if !hasProperties(s) {
			ruleResults.Add(
				"Object must have at least one property",
				s.GetResolvedLocation(),
			)
		}
	}

	recurse.RecurseObjects(s, callback)
	return *ruleResults
}

// hasProperties reports whether s defines any property, directly or through a
// '$ref' -- a schema that is only a '$ref' has the properties of its target.
func hasProperties(s *schema.ExtendedSchema) bool {
	if len(s.Properties)+len(s.PatternProperties) > 0 {
		return true
	}

	if _, ok := s.GetAdditionalProperties().(*schema.ExtendedSchema); ok {
		return true
	}

	ref := s.GetRefSchema()

	return ref != nil && !ref.IsSelfReference() && hasProperties(ref)
}

func (r ObjectsMustHaveProperties) GetSeverity() Severity {
	return SeverityError
}
