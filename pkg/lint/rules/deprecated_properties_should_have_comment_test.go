package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestDeprecatedPropertiesShouldHaveComment(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "no deprecated properties",
			schemaPath:  "testdata/deprecated_properties/no_deprecated_properties.json",
			nViolations: 0,
		},
		{
			name:        "deprecated properties without comment",
			schemaPath:  "testdata/deprecated_properties/deprecated_properties_without_comment.json",
			nViolations: 1,
		},
		{
			name:        "deprecated properties with comment",
			schemaPath:  "testdata/deprecated_properties/deprecated_properties_with_comment.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			rule := DeprecatedPropertiesShouldHaveComment{}
			ruleResults := rule.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults.Violations))
			}
		})
	}
}
