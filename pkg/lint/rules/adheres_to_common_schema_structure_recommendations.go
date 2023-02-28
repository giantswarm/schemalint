package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type AdheresToCommonSchemaStructureRecommendations struct{}

func (r AdheresToCommonSchemaStructureRecommendations) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	recommendedProperties := getRecommendedProperties()

	schemaProperties := schema.GetProperties()
	for _, recommendedProperty := range recommendedProperties {
		if property, ok := schemaProperties[recommendedProperty.Name]; ok {
			if !property.IsType(recommendedProperty.Type) {
				ruleResults.Add(fmt.Sprintf("Root-level property '%s' should be of type '%s'.", recommendedProperty.Name, recommendedProperty.Type))
			}
		} else {
			ruleResults.Add(fmt.Sprintf("Root-level property '%s' should be present.", recommendedProperty.Name))
		}

	}

	return *ruleResults
}

type propertyRecommendation struct {
	Name string
	Type string
}

func getRecommendedProperties() []propertyRecommendation {
	recommendedProperties := []propertyRecommendation{
		{
			Name: "internal",
			Type: "object",
		},
		{
			Name: "providerSpecific",
			Type: "object",
		},
	}

	return recommendedProperties
}

func (r AdheresToCommonSchemaStructureRecommendations) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
