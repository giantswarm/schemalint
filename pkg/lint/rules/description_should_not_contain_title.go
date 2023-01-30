package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionShouldNotContainTitle struct{}

func (r DescriptionShouldNotContainTitle) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereDescriptionsExist()
	for path, annotations := range propertyAnnotationsMap {
		if descriptionContainsTitle(annotations.GetDescription(), annotations.GetTitle()) {
			ruleResults.Add(fmt.Sprintf("Property '%s' description should not repeat the title.", path))
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

func (r DescriptionShouldNotContainTitle) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}