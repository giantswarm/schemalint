package rules

import (
	"github.com/giantswarm/schemalint/pkg/lint/pam"
	"github.com/giantswarm/schemalint/pkg/schema"
)

type TitleMustBeSentenceCase struct{}

func (r TitleMustBeSentenceCase) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s).WhereTitlesExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if !stringStartsCapitalized(annotations.GetTitle()) {
			ruleResults.Add(
				"Property title must start with a capital letter",
				resolvedLocation,
			)
		}
	}

	return *ruleResults
}

func (r TitleMustBeSentenceCase) GetSeverity() Severity {
	return SeverityError
}
