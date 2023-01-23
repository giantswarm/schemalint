package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestTitleExists(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "has no title",
			schemaPath:  "testdata/8_missing_titles.json",
			nViolations: 8,
		},
		{
			name:        "has no title - referenced",
			schemaPath:  "testdata/9_missing_titles_referenced.json",
			nViolations: 9,
		},
		{
			name:        "has title",
			schemaPath:  "testdata/has_titles.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := TitleExists{}
			ruleResults := titleExistsRule.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults.Violations))
			}
		})
	}
}
