package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

func TestTitleExists(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "has no title",
			schemaPath:  "testdata/title_exists/8_missing_titles.json",
			nViolations: 8,
		},
		{
			name:        "has no title - referenced",
			schemaPath:  "testdata/title_exists/9_missing_titles_referenced.json",
			nViolations: 9,
		},
		{
			name:        "referenced has missing titles - override",
			schemaPath:  "testdata/title_exists/9_missing_titles_referenced_overridden.json",
			nViolations: 0,
		},
		{
			name:        "has title",
			schemaPath:  "testdata/title_exists/has_titles.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			titleExistsRule := TitleExists{}
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
