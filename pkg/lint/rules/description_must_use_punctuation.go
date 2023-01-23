package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

const AllowedEndings = ".!?"

type DescriptionMustUsePunctuation struct{}

func (r DescriptionMustUsePunctuation) Verify(schema *schemautils.ExtendedSchema) []string {
	return lint.RecursePropertiesWithDescription(schema, checkDescriptionMustUsePunctuation)
}

func checkDescriptionMustUsePunctuation(schema *schemautils.ExtendedSchema) []string {
	if !endsWithPunctuation(schema.Description) {
		return []string{fmt.Sprintf("Property '%s' description must end with punctuation.", schema.GetConciseLocation())}
	}
	return []string{}
}

func endsWithPunctuation(s string) bool {
	lastChar := rune(s[len(s)-1])
	return strings.ContainsRune(AllowedEndings, lastChar)
}

func (r DescriptionMustUsePunctuation) GetSeverity() lint.Severity {
	return lint.SeverityError
}
