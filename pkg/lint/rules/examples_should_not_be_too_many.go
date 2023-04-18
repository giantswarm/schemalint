package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
	"github.com/giantswarm/schemalint/v2/pkg/lint/utils"
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

type ExamplesShouldNotBeTooMany struct{}

const maxExamples = 5

func (r ExamplesShouldNotBeTooMany) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema)
	for resolvedLocation, propertyAnnotations := range propertyAnnotationsMap {
		examples := propertyAnnotations.GetExamples()
		if len(examples) > maxExamples {
			ruleResults.Add(
				fmt.Sprintf(
					"Property '%s' should not have more than %d examples.",
					schemautils.ConvertToConciseLocation(resolvedLocation),
					maxExamples,
				),
				resolvedLocation,
			)
		}
	}

	return *ruleResults
}

func (r ExamplesShouldNotBeTooMany) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
