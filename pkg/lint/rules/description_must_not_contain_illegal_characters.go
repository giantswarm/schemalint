package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionMustNotContainIllegalCharacters struct{}

var descriptionIllegalCharacters = []string{"\n", "\r", "\t", "  "}

func (r DescriptionMustNotContainIllegalCharacters) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereDescriptionsExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		description := annotations.GetDescription()
		if containedIllegalChars := getIllegalCharacterIn(description, descriptionIllegalCharacters); len(containedIllegalChars) > 0 {
			ruleResults.Add(
				fmt.Sprintf("Property '%s' description must not contain illegal characters: %q", schemautils.ConvertToConciseLocation(resolvedLocation), containedIllegalChars),
				resolvedLocation,
			)
		}
		if containsLeadingOrTrailingSpace(description) {
			ruleResults.Add(
				fmt.Sprintf("Property '%s' description must not contain leading or trailing spaces", schemautils.ConvertToConciseLocation(resolvedLocation)),
				resolvedLocation,
			)
		}
	}
	return *ruleResults
}

func (r DescriptionMustNotContainIllegalCharacters) GetSeverity() lint.Severity {
	return lint.SeverityError
}
