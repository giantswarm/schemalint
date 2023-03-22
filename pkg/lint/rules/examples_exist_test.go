package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestExampleExists(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "has no example",
			schemaPath:  "testdata/examples_exist/no_examples.json",
			nViolations: 1,
		},
		{
			name:        "has example",
			schemaPath:  "testdata/examples_exist/has_examples.json",
			nViolations: 0,
		},
		{
			name:        "no examples, but boolean",
			schemaPath:  "testdata/examples_exist/no_examples_but_boolean.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			exampleExistsRule := ExampleExists{}
			ruleResults := exampleExistsRule.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults.Violations))
			}
		})
	}
}
