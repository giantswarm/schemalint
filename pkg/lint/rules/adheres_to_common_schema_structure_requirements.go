package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type AdheresToCommonSchemaStructureRequirements struct{}

func (r AdheresToCommonSchemaStructureRequirements) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}

	requiredProperties := getRequiredRootProperties()

	rootSchemaProperties := s.GetProperties()
	var globalSchemaProperties map[string]*schema.ExtendedSchema
	if globalSchema, ok := rootSchemaProperties["global"]; ok {
		globalSchemaProperties = globalSchema.GetProperties()
	}

	for _, requiredProperty := range requiredProperties {
		_, requiredPropertyFound := rootSchemaProperties[requiredProperty]
		_, globalRequiredPropertyFound := globalSchemaProperties[requiredProperty]

		if !requiredPropertyFound && !globalRequiredPropertyFound {
			ruleResults.Add(
				"Global or root-level property must be present",
				"/properties/"+requiredProperty,
			)
		}

	}

	allAllowedRootProperties := getAllAllowedRootPropertiesNamesSet()
	for key, s := range rootSchemaProperties {
		if _, ok := allAllowedRootProperties[key]; !ok {
			ruleResults.Add(
				"Root-level property is not allowed",
				s.GetResolvedLocation(),
			)
		}
	}

	if s.AdditionalProperties != false {
		ruleResults.Add(
			"Additional properties must not be allowed at the root level",
			"",
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
		"global", // values shared with subcharts
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
