package rules

import (
	"github.com/giantswarm/schemalint/pkg/lint/pam"
	"github.com/giantswarm/schemalint/pkg/schema"
)

type ExampleExists struct{}

func (r ExampleExists) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}
	propertyAnnotationsMap := pam.BuildPropertyAnnotationsMap(s)
	for resolvedLocation, annotations := range propertyAnnotationsMap {
		if len(annotations.GetExamples()) == 0 && recommendExamplesAt(s, resolvedLocation) {
			ruleResults.Add(
				"Property should provide one or more examples",
				resolvedLocation,
			)
		}
	}
	return *ruleResults
}

func recommendExamplesAt(s *schema.ExtendedSchema, location string) bool {
	schemas := s.GetSchemasAt(location)
	isString := allSchemasAreString(schemas)
	if !isString {
		return false
	}
	isRestricted := oneStringSchemaIsRestricted(schemas)
	return isRestricted
}

func allSchemasAreString(schemas []*schema.ExtendedSchema) bool {
	for _, s := range schemas {
		if !s.IsString() {
			return false
		}
	}
	return true
}

func oneStringSchemaIsRestricted(schemas []*schema.ExtendedSchema) bool {
	for _, s := range schemas {
		if s.Pattern != nil || s.Format != "" {
			return true
		}
	}
	return false
}

func (r ExampleExists) GetSeverity() Severity {
	return SeverityRecommendation
}
