package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/schema"
)

func TestAvoidUnevaluated(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "uses unevaluated properties",
			schemaPath:  "testdata/avoid_unevaluated/uses_unevaluated_properties.json",
			nViolations: 1,
		},
		{
			name:        "uses unevaluated items",
			schemaPath:  "testdata/avoid_unevaluated/uses_unevaluated_items.json",
			nViolations: 1,
		},
		{
			name:        "does not use unevaluated properties or items",
			schemaPath:  "testdata/avoid_unevaluated/correct.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := AvoidUnevaluated{}
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
