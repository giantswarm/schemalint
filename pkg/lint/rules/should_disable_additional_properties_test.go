package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
)

func TestShouldDisableAdditionalProperties(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
		rules       []lint.Rule
	}{
		{
			name:        "additional properties not disabled",
			schemaPath:  "testdata/additional_properties/not_disabled.json",
			nViolations: 1,
			rules:       []lint.Rule{ShouldDisableAdditionalProperties{}},
		},
		{
			name:        "additional properties disabled",
			schemaPath:  "testdata/additional_properties/disabled.json",
			nViolations: 0,
			rules:       []lint.Rule{ShouldDisableAdditionalProperties{}},
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
