package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestMustUseCorrectDialect(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "specifies no dialect",
			schemaPath:  "testdata/dialect/no_dialect.json",
			nViolations: 1,
		},
		{
			name:        "specifies incorrect dialect",
			schemaPath:  "testdata/dialect/incorrect_dialect.json",
			nViolations: 1,
		},
		{
			name:        "specifies correct dialect",
			schemaPath:  "testdata/dialect/correct_dialect.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			mustUseCorrectDialect := MustUseCorrectDialect{}
			ruleResults := mustUseCorrectDialect.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults.Violations))
			}
		})
	}
}
