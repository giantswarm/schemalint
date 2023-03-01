package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestAvoidRecursionKeywords(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "case 0: contains dynamicRef",
			schemaPath:  "testdata/avoid_recursion_keywords/dynamic_ref.json",
			nViolations: 1,
		},
		{
			name:        "case 1: correct",
			schemaPath:  "testdata/avoid_recursion_keywords/correct.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := AvoidRecursionKeywords{}
			ruleResults := titleExistsRule.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults.Violations))
			}
		})
	}
}
