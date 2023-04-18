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
		if _, ok := schemaProperties[recommendedProperty]; !ok {
			ruleResults.Add(
				fmt.Sprintf(
					"Root-level property '%s' should be present.",
					recommendedProperty,
				),
				schema.GetResolvedLocation(),
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

func (r AdheresToCommonSchemaStructureRecommendations) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
