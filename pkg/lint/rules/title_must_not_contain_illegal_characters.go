package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type TitleMustNotContainIllegalCharacters struct{}

var titleIllegalCharacters = []string{"\n", "\r", "\t", "  ", ".", ",", "?", "!"}

func (r TitleMustNotContainIllegalCharacters) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereTitlesExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		title := annotations.GetTitle()
		if containedIllegalChars := getIllegalCharacterIn(title, titleIllegalCharacters); len(containedIllegalChars) > 0 {
			ruleResults.Add(
				fmt.Sprintf(
					"Property '%s' title must not contain illegal characters: %q",
					schemautils.ConvertToConciseLocation(resolvedLocation),
					containedIllegalChars,
				),
				resolvedLocation,
			)
		}
		if containsLeadingOrTrailingSpace(title) {
			ruleResults.Add(fmt.Sprintf("Property '%s' title must not contain leading or trailing spaces.", resolvedLocation), schema.GetResolvedLocation())
		}
	}
	return *ruleResults
}

func (r TitleMustNotContainIllegalCharacters) GetSeverity() lint.Severity {
	return lint.SeverityError
}
