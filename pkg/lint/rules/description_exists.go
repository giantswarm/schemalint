package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/pam"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type DescriptionExists struct{}

func (r DescriptionExists) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s)

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if annotations.GetDescription() == "" {
			ruleResults.Add(
				"Property should have a description",
				resolvedLocation,
			)
		}
	}

	return *ruleResults
}

func (r DescriptionExists) GetSeverity() Severity {
	return SeverityRecommendation
}
