package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/lint/pam"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type DescriptionMustBeSentenceCase struct{}

func (r DescriptionMustBeSentenceCase) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s).WhereDescriptionsExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if !stringStartsCapitalized(annotations.GetDescription()) {
			ruleResults.Add(
				"Property description must start with a capital letter",
				resolvedLocation,
			)
		}
	}

	return *ruleResults
}

func (r DescriptionMustBeSentenceCase) GetSeverity() Severity {
	return SeverityError
}
