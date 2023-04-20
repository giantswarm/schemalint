package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type AdheresToCommonSchemaStructureRequirements struct{}

func (r AdheresToCommonSchemaStructureRequirements) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}

	requiredProperties := getRequiredRootProperties()

	schemaProperties := s.GetProperties()
	for _, requiredProperty := range requiredProperties {
		if _, ok := schemaProperties[requiredProperty]; !ok {
			ruleResults.Add(
				fmt.Sprintf("Root-level property %s must be present", requiredProperty),
				s.GetResolvedLocation(),
			)
		}

	}

	allAllowedRootProperties := getAllAllowedRootPropertiesNamesSet()
	for key, s := range schemaProperties {
		if _, ok := allAllowedRootProperties[key]; !ok {
			ruleResults.Add(
				fmt.Sprintf("Root-level property %s is not allowed", key),
				s.GetResolvedLocation(),
			)
		}
	}

	if s.AdditionalProperties != false {
		ruleResults.Add(
			"Additional properties must not be allowed at the root level",
			s.GetResolvedLocation(),
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
		"kubectlImage",
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

func (r AdheresToCommonSchemaStructureRequirements) GetSeverity() Severity {
	return SeverityError
}
