package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
)

func TestArrayItemsMustHaveSingleType(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
		rules       []lint.Rule
	}{
		{
			name:        "case 0: array with 'contains'",
			schemaPath:  "testdata/array_items_must_have_single_type/contains.json",
			nViolations: 1,
			rules:       []lint.Rule{ArrayItemsMustHaveSingleType{}},
		},
		{
			name:        "case 1: array with 'prefixItems'",
			schemaPath:  "testdata/array_items_must_have_single_type/prefix_items.json",
			nViolations: 1,
			rules:       []lint.Rule{ArrayItemsMustHaveSingleType{}},
		},
		{
			name:        "case 1: array with 'prefixItems' and 'contains'",
			schemaPath:  "testdata/array_items_must_have_single_type/multiple_illegal_keywords.json",
			nViolations: 1,
			rules:       []lint.Rule{ArrayItemsMustHaveSingleType{}},
		},
		{
			name:        "case 2: correct array",
			schemaPath:  "testdata/array_items_must_have_single_type/correct.json",
			nViolations: 0,
			rules:       []lint.Rule{ArrayItemsMustHaveSingleType{}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			ruleResults := []lint.RuleViolation{}
			for _, rule := range tc.rules {
				ruleResults = append(ruleResults, rule.Verify(schema).Violations...)
			}

			if len(ruleResults) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults))
			}
		})
	}
}
