package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type AdheresToCommonSchemaStructureRecommendations struct{}

func (r AdheresToCommonSchemaStructureRecommendations) Verify(
	s *schema.ExtendedSchema,
) RuleResults {
	ruleResults := &RuleResults{}

	recommendedProperties := getRecommendedRootProperties()

	schemaProperties := s.GetProperties()
	for _, recommendedProperty := range recommendedProperties {
		if _, ok := schemaProperties[recommendedProperty]; !ok {
			ruleResults.Add(
				fmt.Sprintf(
					"Root-level property %s should be present",
					recommendedProperty,
				),
				s.GetResolvedLocation(),
			)
		}

	}

	return *ruleResults
}

func getRecommendedRootProperties() []string {
	recommendedProperties := []string{
		"internal",
		"providerSpecific",
	}

	return recommendedProperties
}

func (r AdheresToCommonSchemaStructureRecommendations) GetSeverity() Severity {
	return SeverityRecommendation
}
