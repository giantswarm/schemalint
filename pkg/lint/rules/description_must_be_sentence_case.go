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

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereDescriptionsExist()

	for path, annotations := range propertyAnnotationsMap {
		if !descriptionStartsCapitalized(annotations.GetDescription()) {
			ruleResults.Add(fmt.Sprintf("Property '%s' description must start with a capital letter.", path))
		}
	}

	return *ruleResults
}

func descriptionStartsCapitalized(description string) bool {
	return unicode.IsUpper(rune(description[0]))

}

func (r DescriptionMustBeSentenceCase) GetSeverity() lint.Severity {
	return lint.SeverityError
}
