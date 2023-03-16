package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
)

func TestExamplesShouldNotBeTooMany(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "too many examples",
			schemaPath:  "testdata/examples_correct/too_many_examples.json",
			nViolations: 1,
		},
		{
			name:        "correct number of examples",
			schemaPath:  "testdata/examples_correct/correct.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := ExamplesShouldNotBeTooMany{}
			ruleResults := titleExistsRule.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults.Violations))
			}
		})
	}
}
