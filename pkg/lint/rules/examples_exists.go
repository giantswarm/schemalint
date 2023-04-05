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
		if len(annotations.GetExamples()) == 0 && recommendExamplesAt(schema, resolvedLocation) {
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

func recommendExamplesAt(schema *schemautils.ExtendedSchema, location string) bool {
	schemas := schema.GetSchemasAt(location)
	isString := allSchemasAreString(schemas)
	if !isString {
		return false
	}
	isRestricted := oneStringSchemaIsRestricted(schemas)
	return isRestricted
}

func allSchemasAreString(schemas []*schemautils.ExtendedSchema) bool {
	for _, schema := range schemas {
		if !schema.IsString() {
			return false
		}
	}
	return true
}

func oneStringSchemaIsRestricted(schemas []*schemautils.ExtendedSchema) bool {
	for _, schema := range schemas {
		if schema.Pattern != nil || schema.Format != "" {
			return true
		}
	}
	return false
}

func (r ExampleExists) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
