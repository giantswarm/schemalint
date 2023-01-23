package rules

import (
	"fmt"
	"unicode"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionMustBeSentenceCase struct{}

func (r DescriptionMustBeSentenceCase) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	callback := func(schema *schemautils.ExtendedSchema) {
		checkDescriptionMustBeSentenceCase(schema, ruleResults)
	}
	utils.RecursePropertiesWithDescription(schema, callback)
	return *ruleResults
}

func checkDescriptionMustBeSentenceCase(schema *schemautils.ExtendedSchema, ruleResults *lint.RuleResults) {
	if !unicode.IsUpper(rune(schema.Description[0])) {
		ruleResults.Add(fmt.Sprintf("Property '%s' description must start with a capital letter.", schema.GetConciseLocation()))
	}
}

func (r DescriptionMustBeSentenceCase) GetSeverity() lint.Severity {
	return lint.SeverityError
}
