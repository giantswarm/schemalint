package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestDescriptionExists(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "has no description",
			schemaPath:  "testdata/8_missing_descriptions.json",
			nViolations: 8,
		},
		{
			name:        "has description",
			schemaPath:  "testdata/has_descriptions.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		schema, err := lint.Compile(tc.schemaPath)
		if err != nil {
			t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
		}
		descriptionExistsRule := DescriptionExists{}
		ruleViolations := descriptionExistsRule.Verify(schema)

		if len(ruleViolations) != tc.nViolations {
			t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleViolations))
		}
	}
}