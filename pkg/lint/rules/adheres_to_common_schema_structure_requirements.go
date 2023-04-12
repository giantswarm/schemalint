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
		if _, ok := schemaProperties[requiredProperty]; !ok {
			ruleResults.Add(
				fmt.Sprintf("Root-level property '%s' must be present.", requiredProperty),
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

	if schema.AdditionalProperties != false {
		ruleResults.Add(
			"Additional properties must not be allowed at the root level.",
			schema.GetResolvedLocation(),
		)
	}

	return *ruleResults
}

func getRequiredRootProperties() []string {
	requiredProperties := []string{
		"metadata",
		"connectivity",
		"controlPlane",
		"nodePools",
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

	allAllowedRootProperties = append(
		allAllowedRootProperties,
		requireRootProperties...,
	)
	allAllowedRootProperties = append(
		allAllowedRootProperties,
		recommendedRootProperties...,
	)

	allAllowedRootPropertiesMap := make(map[string]bool)
	for _, property := range allAllowedRootProperties {
		allAllowedRootPropertiesMap[property] = true
	}

	return allAllowedRootPropertiesMap
}

func (r AdheresToCommonSchemaStructureRequirements) GetSeverity() lint.Severity {
	return lint.SeverityError
}
