package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

func TestNumbersShouldBeConstrained(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "unconstrained number",
			schemaPath:  "testdata/numbers_should_be_constrained/unconstrained_number.json",
			nViolations: 2,
		},
		{
			name:        "constrained number",
			schemaPath:  "testdata/numbers_should_be_constrained/correct.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := NumbersShouldBeConstrained{}
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
