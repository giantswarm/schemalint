package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type TitleShouldNotContainParentsTitle struct{}

func (r TitleShouldNotContainParentsTitle) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereTitlesExist()

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
				fmt.Sprintf(
					"Property '%s' title should not contain the parent's title '%s'.",
					schemautils.ConvertToConciseLocation(resolvedLocation),
					parentTitle,
				),
				resolvedLocation,
			)
		}
	}
	return *ruleResults
}

func (r TitleShouldNotContainParentsTitle) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
