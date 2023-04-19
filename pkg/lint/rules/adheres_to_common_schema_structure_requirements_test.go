package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/schema"
)

func TestAdheheresToCommonSchemaStructureRequirements(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "case 0: missing metadata",
			schemaPath:  "testdata/common_schema_structure_requirements/missing_metadata.json",
			nViolations: 1,
		},
		{
			name:        "case 1: missing connectivity",
			schemaPath:  "testdata/common_schema_structure_requirements/missing_connectivity.json",
			nViolations: 1,
		},
		{
			name:        "case 2: missing controlPlane",
			schemaPath:  "testdata/common_schema_structure_requirements/missing_control_plane.json",
			nViolations: 1,
		},
		{
			name:        "case 3: missing nodePools",
			schemaPath:  "testdata/common_schema_structure_requirements/missing_node_pools.json",
			nViolations: 1,
		},
		{
			name:        "case 4: missing all",
			schemaPath:  "testdata/common_schema_structure_requirements/missing_all.json",
			nViolations: 4,
		},
		{
			name:        "case 5: correct",
			schemaPath:  "testdata/common_schema_structure_requirements/correct.json",
			nViolations: 0,
		},
		{
			name:        "case 6: correct with optional fields",
			schemaPath:  "testdata/common_schema_structure_requirements/correct_with_optional_fields.json",
			nViolations: 0,
		},
		{
			name:        "case 7: too many",
			schemaPath:  "testdata/common_schema_structure_requirements/too_many.json",
			nViolations: 1,
		},
		{
			name:        "case 8: additional properties not set to false",
			schemaPath:  "testdata/common_schema_structure_requirements/additional_properties_not_set_to_false.json",
			nViolations: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := AdheresToCommonSchemaStructureRequirements{}
			ruleResults := titleExistsRule.Verify(schema)

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
