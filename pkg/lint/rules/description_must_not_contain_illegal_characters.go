package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint/pam"
	"github.com/giantswarm/schemalint/pkg/schema"
)

type DescriptionMustNotContainIllegalCharacters struct{}

var descriptionIllegalCharacters = []string{"\n", "\r", "\t", "  "}

func (r DescriptionMustNotContainIllegalCharacters) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s).WhereDescriptionsExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		description := annotations.GetDescription()
		if containedIllegalChars := getIllegalCharacterIn(description, descriptionIllegalCharacters); len(
			containedIllegalChars,
		) > 0 {
			ruleResults.Add(
				fmt.Sprintf(
					"Property description must not contain illegal characters: %q",
					descriptionIllegalCharacters,
				),
				resolvedLocation,
			)
		}
		if containsLeadingOrTrailingSpace(description) {
			ruleResults.Add(
				"Property description must not contain leading or trailing spaces",
				resolvedLocation,
			)
		}
	}
	return *ruleResults
}

func (r DescriptionMustNotContainIllegalCharacters) GetSeverity() Severity {
	return SeverityError
}
