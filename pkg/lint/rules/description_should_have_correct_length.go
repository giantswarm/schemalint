package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint/pam"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

const maxDescriptionLength = 200
const minDescriptionLength = 50

type DescriptionShouldHaveCorrectLength struct{}

func (r DescriptionShouldHaveCorrectLength) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s).WhereDescriptionsExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if descriptionIsTooShort(annotations.GetDescription()) {
			ruleResults.Add(
				fmt.Sprintf(
					"Property description should be more than %d characters",
					minDescriptionLength,
				),
				resolvedLocation,
			)
		}

		if descriptionIsTooLong(annotations.GetDescription()) {
			ruleResults.Add(
				fmt.Sprintf(
					"Property description should be less than %d characters",
					maxDescriptionLength,
				),
				resolvedLocation,
			)

		}
	}

	return *ruleResults
}

func descriptionIsTooLong(description string) bool {
	return len(description) > maxDescriptionLength
}

func descriptionIsTooShort(description string) bool {
	return len(description) < minDescriptionLength
}

func (r DescriptionShouldHaveCorrectLength) GetSeverity() Severity {
	return SeverityRecommendation
}
