package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
)

func TestArraysMustHaveItems(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
		rules       []lint.Rule
	}{
		{
			name:        "array without items",
			schemaPath:  "testdata/arrays_must_have_items/without_items.json",
			nViolations: 1,
		},
		{
			name:        "array with items",
			schemaPath:  "testdata/arrays_must_have_items/with_items.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			rule := ArraysMustHaveItems{}
			ruleResults := rule.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults.Violations))
			}
		})
	}
}
