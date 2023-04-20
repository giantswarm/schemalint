package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint/pam"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type TitleMustNotContainIllegalCharacters struct{}

var titleIllegalCharacters = []string{"\n", "\r", "\t", "  ", ".", ",", "?", "!"}

func (r TitleMustNotContainIllegalCharacters) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s).WhereTitlesExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		title := annotations.GetTitle()
		if containedIllegalChars := getIllegalCharacterIn(title, titleIllegalCharacters); len(
			containedIllegalChars,
		) > 0 {
			ruleResults.Add(
				fmt.Sprintf(
					"Property title must not contain illegal characters: %q",
					titleIllegalCharacters,
				),
				resolvedLocation,
			)
		}
		if containsLeadingOrTrailingSpace(title) {
			ruleResults.Add(
				"Property title must not contain leading or trailing spaces.",
				s.GetResolvedLocation(),
			)
		}
	}
	return *ruleResults
}

func (r TitleMustNotContainIllegalCharacters) GetSeverity() Severity {
	return SeverityError
}
