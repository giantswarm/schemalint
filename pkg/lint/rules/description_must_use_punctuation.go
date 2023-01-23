package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

const AllowedEndings = ".!?"

type DescriptionMustUsePunctuation struct{}

func (r DescriptionMustUsePunctuation) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	callback := func(schema *schemautils.ExtendedSchema) {
		checkDescriptionMustUsePunctuation(schema, ruleResults)
	}
	utils.RecursePropertiesWithDescription(schema, callback)
	return *ruleResults
}

func checkDescriptionMustUsePunctuation(schema *schemautils.ExtendedSchema, ruleResults *lint.RuleResults) {
	if !endsWithPunctuation(schema.Description) {
		ruleResults.Add(fmt.Sprintf("Property '%s' description must end with punctuation.", schema.GetConciseLocation()))
	}
}

func endsWithPunctuation(s string) bool {
	lastChar := rune(s[len(s)-1])
	return strings.ContainsRune(AllowedEndings, lastChar)
}

func (r DescriptionMustUsePunctuation) GetSeverity() lint.Severity {
	return lint.SeverityError
}
