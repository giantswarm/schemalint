package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

func TestAvoidLogicalContructs(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "case 0: uses if, then, else",
			schemaPath:  "testdata/avoid_logical_constructs/incorrect.json",
			nViolations: 1,
		},
		{
			name:        "case 1: does not use if, then, else",
			schemaPath:  "testdata/avoid_logical_constructs/correct.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			rule := AvoidLogicalConstruct{}
			ruleResults := rule.Verify(s)

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
