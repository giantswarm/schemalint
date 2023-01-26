package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionMustNotContainIllegalCharacters struct{}

var descriptionIllegalCharacters = []string{"\n", "\r", "\t", "  "}

func (r DescriptionMustNotContainIllegalCharacters) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereDescriptionsExist()

	for path, annotations := range propertyAnnotationsMap {
		description := annotations.GetDescription()
		if isInvalid, containedIllegalChars := containsIllegalCharacter(description); isInvalid {
			ruleResults.Add(fmt.Sprintf("Property '%s' description must not contain illegal characters: %s", path, containedIllegalChars))
		}
		if containsLeadingOrTrailingSpace(description) {
			ruleResults.Add(fmt.Sprintf("Property '%s' description must not contain leading or trailing spaces", path))
		}
	}
	return *ruleResults
}

func containsIllegalCharacter(s string) (contains bool, containedIllegalCharacters []string) {
	for _, illegalCharacter := range descriptionIllegalCharacters {
		if strings.Contains(s, illegalCharacter) {
			containedIllegalCharacters = append(containedIllegalCharacters, illegalCharacter)
		}
	}

	return len(containedIllegalCharacters) > 0, containedIllegalCharacters
}

func containsLeadingOrTrailingSpace(s string) bool {
	return strings.HasPrefix(s, " ") || strings.HasSuffix(s, " ")
}

func (r DescriptionMustNotContainIllegalCharacters) GetSeverity() lint.Severity {
	return lint.SeverityError
}
