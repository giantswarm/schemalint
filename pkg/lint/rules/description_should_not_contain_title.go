package rules

import (
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint/pam"
	"github.com/giantswarm/schemalint/pkg/schema"
)

type DescriptionShouldNotContainTitle struct{}

func (r DescriptionShouldNotContainTitle) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s).WhereDescriptionsExist()
	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if descriptionContainsTitle(annotations.GetDescription(), annotations.GetTitle()) {
			ruleResults.Add(
				"Property description should not repeat the title",
				resolvedLocation,
			)
		}
	}
	return *ruleResults
}

func descriptionContainsTitle(description string, title string) bool {
	if title == "" {
		return false
	}

	return strings.Contains(strings.ToLower(description), strings.ToLower(title))
}

func (r DescriptionShouldNotContainTitle) GetSeverity() Severity {
	return SeverityRecommendation
}
