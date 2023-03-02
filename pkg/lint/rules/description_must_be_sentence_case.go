package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionMustBeSentenceCase struct{}

func (r DescriptionMustBeSentenceCase) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereDescriptionsExist()

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if !stringStartsCapitalized(annotations.GetDescription()) {
			ruleResults.Add(
				fmt.Sprintf("Property '%s' description must start with a capital letter.", schemautils.ConvertToConciseLocation(resolvedLocation)),
				resolvedLocation,
			)
		}
	}

	return *ruleResults
}

func (r DescriptionMustBeSentenceCase) GetSeverity() lint.Severity {
	return lint.SeverityError
}
