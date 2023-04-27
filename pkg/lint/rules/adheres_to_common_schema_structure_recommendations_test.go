package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

func TestAdheheresToCommonSchemaStructureRecommendations(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "case 0: missing internal",
			schemaPath:  "testdata/common_schema_structure_recommendations/missing_internal.json",
			nViolations: 1,
		},
		{
			name:        "case 1: missing providerSpecific",
			schemaPath:  "testdata/common_schema_structure_recommendations/missing_provider_specific.json",
			nViolations: 1,
		},
		{
			name:        "case 2: missing all",
			schemaPath:  "testdata/common_schema_structure_recommendations/missing_all.json",
			nViolations: 2,
		},
		{
			name:        "case 3: correct",
			schemaPath:  "testdata/common_schema_structure_recommendations/correct.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := AdheresToCommonSchemaStructureRecommendations{}
			ruleResults := titleExistsRule.Verify(s)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf(
					"Unexpected number of rule violations in test case '%s': Expected %d, got %d",
					tc.name,
					tc.nViolations,
					len(ruleResults.Violations),
				)
			}
		})
	}
}
