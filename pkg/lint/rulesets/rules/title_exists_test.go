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
			name:        "has title",
			schemaPath:  "testdata/has_title.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		schema, err := lint.Compile(tc.schemaPath)
		if err != nil {
			t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
		}
		titleExistsRule := TitleExists{}
		ruleViolations := titleExistsRule.Verify(schema)

		if len(ruleViolations) != tc.nViolations {
			t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleViolations))
		}
	}
}
