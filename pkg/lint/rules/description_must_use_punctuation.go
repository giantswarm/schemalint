package rules

import (
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint/pam"
	"github.com/giantswarm/schemalint/pkg/schema"
)

const AllowedEndings = ".!?"

type DescriptionMustUsePunctuation struct{}

func (r DescriptionMustUsePunctuation) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s).WhereDescriptionsExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if !endsWithPunctuation(annotations.GetDescription()) {
			ruleResults.Add(
				"Property description must end with punctuation",
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

func (r DescriptionMustUsePunctuation) GetSeverity() Severity {
	return SeverityError
}
