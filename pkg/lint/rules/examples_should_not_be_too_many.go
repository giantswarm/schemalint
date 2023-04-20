package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint/pam"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type ExamplesShouldNotBeTooMany struct{}

const maxExamples = 5

func (r ExamplesShouldNotBeTooMany) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s)
	for resolvedLocation, propertyAnnotations := range propertyAnnotationsMap {
		examples := propertyAnnotations.GetExamples()
		if len(examples) > maxExamples {
			ruleResults.Add(
				fmt.Sprintf(
					"Property should not have more than %d examples.",
					maxExamples,
				),
				resolvedLocation,
			)
		}
	}

	return *ruleResults
}

func (r ExamplesShouldNotBeTooMany) GetSeverity() Severity {
	return SeverityRecommendation
}
