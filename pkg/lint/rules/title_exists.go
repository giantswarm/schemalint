package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
	"github.com/giantswarm/schemalint/v2/pkg/lint/utils"
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

type TitleExists struct{}

func (r TitleExists) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema)
	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if annotations.GetTitle() == "" {
			ruleResults.Add(
				fmt.Sprintf(
					"Property '%s' must have a title.",
					schemautils.ConvertToConciseLocation(resolvedLocation)),
				resolvedLocation,
			)
		}
	}
	return *ruleResults
}

func (r TitleExists) GetSeverity() lint.Severity {
	return lint.SeverityError
}
