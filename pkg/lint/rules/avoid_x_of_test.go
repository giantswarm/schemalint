package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

func TestShouldAvoidXOf(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "uses anyOf",
			schemaPath:  "testdata/avoid_x_of/any_of.json",
			nViolations: 1,
		},
		{
			name:        "uses oneOf",
			schemaPath:  "testdata/avoid_x_of/one_of.json",
			nViolations: 1,
		},
		{
			name:        "uses anyOf and oneOf",
			schemaPath:  "testdata/avoid_x_of/any_of_and_one_of.json",
			nViolations: 2,
		},
		{
			name:        "does not use anyOf or oneOf",
			schemaPath:  "testdata/avoid_x_of/correct.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := AvoidXOf{}
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
