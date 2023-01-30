package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type TitleMustBeSentenceCase struct{}

func (r TitleMustBeSentenceCase) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereTitlesExist()

	for path, annotations := range propertyAnnotationsMap {
		if !stringStartsCapitalized(annotations.GetTitle()) {
			ruleResults.Add(fmt.Sprintf("Property '%s' title must start with a capital letter.", path))
		}
	}

	return *ruleResults
}

func (r TitleMustBeSentenceCase) GetSeverity() lint.Severity {
	return lint.SeverityError
}
