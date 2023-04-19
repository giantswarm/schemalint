package rules

import (
	"github.com/giantswarm/schemalint/pkg/lint/pam"
	"github.com/giantswarm/schemalint/pkg/schema"
)

type TitleExists struct{}

func (r TitleExists) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}
	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s)
	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if annotations.GetTitle() == "" {
			ruleResults.Add(
				"Property must have a title",
				resolvedLocation,
			)
		}
	}
	return *ruleResults
}

func (r TitleExists) GetSeverity() Severity {
	return SeverityError
}
