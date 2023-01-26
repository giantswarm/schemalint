package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

const maxDescriptionLength = 200
const minDescriptionLength = 50

type DescriptionShouldHaveCorrectLength struct{}

func (r DescriptionShouldHaveCorrectLength) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereDescriptionsExist()

	for path, annotations := range propertyAnnotationsMap {
		if descriptionIsTooShort(annotations.GetDescription()) {
			ruleResults.Add(fmt.Sprintf("Property '%s' description should be more than %d characters.", path, minDescriptionLength))
		}

		if descriptionIsTooLong(annotations.GetDescription()) {
			ruleResults.Add(fmt.Sprintf("Property '%s' description should be less than %d characters.", path, maxDescriptionLength))
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

func (r DescriptionShouldHaveCorrectLength) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
