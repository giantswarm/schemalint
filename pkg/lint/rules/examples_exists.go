package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type ExampleExists struct{}

func (r ExampleExists) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema)
	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if len(annotations.GetExamples()) == 0 && !annotationBelongsToBoolean(schema, resolvedLocation) {
			ruleResults.Add(
				fmt.Sprintf(
					"Property '%s' should provide one or more examples.",
					schemautils.ConvertToConciseLocation(resolvedLocation),
				),
				resolvedLocation,
			)
		}
	}
	return *ruleResults
}

func annotationBelongsToBoolean(schema *schemautils.ExtendedSchema, location string) bool {
	schemas := schema.GetSchemasAt(location)
	for _, s := range schemas {
		if s.IsBoolean() {
			return true
		}
	}
	return false
}

func (r ExampleExists) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
