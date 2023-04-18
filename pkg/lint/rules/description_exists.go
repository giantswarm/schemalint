package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
	"github.com/giantswarm/schemalint/v2/pkg/lint/utils"
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

type DescriptionExists struct{}

func (r DescriptionExists) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema)

	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if annotations.GetDescription() == "" {
			ruleResults.Add(
				fmt.Sprintf("Property '%s' should have a description.", schemautils.ConvertToConciseLocation(resolvedLocation)),
				resolvedLocation,
			)
		}
	}

	return *ruleResults
}

func (r DescriptionExists) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
