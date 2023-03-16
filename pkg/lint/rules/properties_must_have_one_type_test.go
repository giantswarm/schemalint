package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
)

func TestPropertiesMustHaveOneType(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "property with multiple types",
			schemaPath:  "testdata/properties_must_have_one_type/multiple_types.json",
			nViolations: 1,
		},
		{
			name:        "property with no types",
			schemaPath:  "testdata/properties_must_have_one_type/no_type.json",
			nViolations: 1,
		},
		{
			name:        "property with one type",
			schemaPath:  "testdata/properties_must_have_one_type/one_type.json",
			nViolations: 0,
		},
		{
			name:        "property with no type but ref",
			schemaPath:  "testdata/properties_must_have_one_type/no_type_reference.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			exampleExistsRule := PropertiesMustHaveOneType{}
			ruleResults := exampleExistsRule.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults.Violations))
			}
		})
	}
}
