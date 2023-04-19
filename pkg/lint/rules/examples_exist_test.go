package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/schema"
)

func TestExampleExists(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "restricted string and no examples",
			schemaPath:  "testdata/examples_exist/no_examples.json",
			nViolations: 2,
		},
		{
			name:        "restricted string and examples",
			schemaPath:  "testdata/examples_exist/has_examples.json",
			nViolations: 0,
		},
		{
			name:        "no examples, but boolean",
			schemaPath:  "testdata/examples_exist/no_examples_but_boolean.json",
			nViolations: 0,
		},
		{
			name:        "no examples, but integer",
			schemaPath:  "testdata/examples_exist/no_examples_but_integer.json",
			nViolations: 0,
		},
		{
			name:        "no examples and string but not restricted",
			schemaPath:  "testdata/examples_exist/no_examples_and_string_but_not_restricted.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			exampleExistsRule := ExampleExists{}
			ruleResults := exampleExistsRule.Verify(s)

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
