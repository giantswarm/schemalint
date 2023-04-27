package rules

import (
	"strings"

	"github.com/giantswarm/schemalint/v2/pkg/lint/pam"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type TitleShouldNotContainParentsTitle struct{}

func (r TitleShouldNotContainParentsTitle) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s).WhereTitlesExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		title := annotations.GetTitle()
		parentAnnotations, err := propertyAnnotationsMap.GetParentAnnotations(resolvedLocation)
		if err != nil {
			continue
		}
		parentTitle := parentAnnotations.GetTitle()

		if parentTitle == "" {
			continue
		}

		if strings.Contains(strings.ToLower(title), strings.ToLower(parentTitle)) {
			ruleResults.Add(
				"Property title should not contain the parent's title",
				resolvedLocation,
			)
		}
	}
	return *ruleResults
}

func (r TitleShouldNotContainParentsTitle) GetSeverity() Severity {
	return SeverityRecommendation
}
