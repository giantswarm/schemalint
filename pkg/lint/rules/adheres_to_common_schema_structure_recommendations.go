package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type AdheresToCommonSchemaStructureRecommendations struct{}

func (r AdheresToCommonSchemaStructureRecommendations) Verify(
	schema *schemautils.ExtendedSchema,
) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	recommendedProperties := getRecommendedRootProperties()

	schemaProperties := schema.GetProperties()
	for _, recommendedProperty := range recommendedProperties {
		property, ok := schemaProperties[recommendedProperty.Name]
		if ok && !property.IsType(recommendedProperty.Type) {
			ruleResults.Add(
				fmt.Sprintf(
					"Root-level property '%s' should be of type '%s'.",
					recommendedProperty.Name,
					recommendedProperty.Type,
				),
				property.GetResolvedLocation(),
			)
		}

		if !ok {
			ruleResults.Add(
				fmt.Sprintf(
					"Root-level property '%s' should be present.",
					recommendedProperty.Name,
				),
				schema.GetResolvedLocation(),
			)
		}

	}

	return *ruleResults
}

func getRecommendedRootProperties() []propertyNameWithType {
	recommendedProperties := []propertyNameWithType{
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
