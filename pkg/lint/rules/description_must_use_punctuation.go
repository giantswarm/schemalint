package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
	"github.com/giantswarm/schemalint/v2/pkg/lint/utils"
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

const AllowedEndings = ".!?"

type DescriptionMustUsePunctuation struct{}

func (r DescriptionMustUsePunctuation) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereDescriptionsExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if !endsWithPunctuation(annotations.GetDescription()) {
			ruleResults.Add(
				fmt.Sprintf("Property '%s' description must end with punctuation.", schemautils.ConvertToConciseLocation(resolvedLocation)),
				resolvedLocation,
			)
		}
	}

	return *ruleResults
}

func endsWithPunctuation(s string) bool {
	lastChar := rune(s[len(s)-1])
	return strings.ContainsRune(AllowedEndings, lastChar)
}

func (r DescriptionMustUsePunctuation) GetSeverity() lint.Severity {
	return lint.SeverityError
}
