package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionMustNotContainIllegalCharacters struct{}

var descriptionIllegalCharacters = []string{"\n", "\r", "\t", "  "}

func (r DescriptionMustNotContainIllegalCharacters) Verify(schema *schemautils.ExtendedSchema) []string {
	return lint.RecursePropertiesWithDescription(schema, checkDescriptionDoesNotContainIllegalCharacters)
}

func checkDescriptionDoesNotContainIllegalCharacters(schema *schemautils.ExtendedSchema) []string {
	ruleViolations := []string{}

	isInvalid, containedIllegalChars := containsIllegalCharacter(schema.Description)
	if isInvalid {
		ruleViolations = append(ruleViolations, fmt.Sprintf(
			"Property '%s' description must not contain %s",
			schema.GetConciseLocation(),
			strings.Join(containedIllegalChars, "', '"),
		))
	}

	if containsLeadingOrTrailingSpace(schema.Description) {
		ruleViolations = append(ruleViolations, fmt.Sprintf(
			"Property '%s' description must not contain leading or trailing spaces",
			schema.GetConciseLocation(),
		))
	}

	return ruleViolations
}

func containsIllegalCharacter(s string) (contains bool, containedIllegalCharacters []string) {
	contains = false

	for _, illegalCharacter := range descriptionIllegalCharacters {
		if strings.Contains(s, illegalCharacter) {
			contains = true
			containedIllegalCharacters = append(containedIllegalCharacters, illegalCharacter)
		}
	}

	return contains, containedIllegalCharacters
}

func containsLeadingOrTrailingSpace(s string) bool {
	return strings.HasPrefix(s, " ") || strings.HasSuffix(s, " ")
}

func (r DescriptionMustNotContainIllegalCharacters) GetSeverity() lint.Severity {
	return lint.SeverityError
}
