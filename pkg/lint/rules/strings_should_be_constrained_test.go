package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestStringsShouldBeConstrained(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "unconstrained string",
			schemaPath:  "testdata/strings_should_be_constrained/unconstrained_string.json",
			nViolations: 1,
		},
		{
			name:        "constrained string",
			schemaPath:  "testdata/strings_should_be_constrained/correct.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := StringsShouldBeConstrained{}
			ruleResults := titleExistsRule.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults.Violations))
			}
		})
	}
}
