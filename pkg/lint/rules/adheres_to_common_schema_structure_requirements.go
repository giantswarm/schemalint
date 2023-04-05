package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type AdheresToCommonSchemaStructureRequirements struct{}

func (r AdheresToCommonSchemaStructureRequirements) Verify(
	schema *schemautils.ExtendedSchema,
) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	requiredProperties := getRequiredRootProperties()

	schemaProperties := schema.GetProperties()
	for _, requiredProperty := range requiredProperties {
		property, ok := schemaProperties[requiredProperty.Name]
		if ok && !property.IsType(requiredProperty.Type) {
			ruleResults.Add(
				fmt.Sprintf(
					"Root-level property '%s' must be of type '%s'.",
					requiredProperty.Name,
					requiredProperty.Type,
				),
				property.GetResolvedLocation(),
			)
		}

		if !ok {
			ruleResults.Add(
				fmt.Sprintf("Root-level property '%s' must be present.", requiredProperty.Name),
				schema.GetResolvedLocation(),
			)
		}

	}

	allAllowedRootProperties := getAllAllowedRootPropertiesNamesSet()
	for key, schema := range schemaProperties {
		if _, ok := allAllowedRootProperties[key]; !ok {
			ruleResults.Add(
				fmt.Sprintf("Root-level property '%s' is not allowed.", key),
				schema.GetResolvedLocation(),
			)
		}
	}

	return *ruleResults
}

type propertyNameWithType struct {
	Name string
	Type string
}

func getRequiredRootProperties() []propertyNameWithType {
	requiredProperties := []propertyNameWithType{
		{
			Name: "metadata",
			Type: "object",
		},
		{
			Name: "connectivity",
			Type: "object",
		},
		{
			Name: "controlPlane",
			Type: "object",
		},
		{
			Name: "nodePools",
			Type: "array",
		},
	}

	return requiredProperties
}

func getAddtionalAllowedRootPropertiesNames() []string {
	return []string{
		"managementCluster",
		"baseDomain",
		"provider",
		"cluster-shared",
	}
}

func getAllAllowedRootPropertiesNamesSet() map[string]bool {
	requireRootProperties := getRequiredRootProperties()
	recommendedRootProperties := getRecommendedRootProperties()

	allAllowedRootProperties := getAddtionalAllowedRootPropertiesNames()

	for _, requiredProperty := range requireRootProperties {
		allAllowedRootProperties = append(allAllowedRootProperties, requiredProperty.Name)
	}
	for _, recommendedProperty := range recommendedRootProperties {
		allAllowedRootProperties = append(allAllowedRootProperties, recommendedProperty.Name)
	}
	allAllowedRootPropertiesMap := make(map[string]bool)
	for _, property := range allAllowedRootProperties {
		allAllowedRootPropertiesMap[property] = true
	}

	return allAllowedRootPropertiesMap
}

func (r AdheresToCommonSchemaStructureRequirements) GetSeverity() lint.Severity {
	return lint.SeverityError
}
