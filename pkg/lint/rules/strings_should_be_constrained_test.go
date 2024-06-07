package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

func TestStringsShouldBeConstrained(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "unconstrained string",
			schemaPath:  "testdata/strings_should_be_constrained/unconstrained_string.json",
			nViolations: 1,
		},
		{
			name:        "constrained with const",
			schemaPath:  "testdata/strings_should_be_constrained/correct_const.json",
			nViolations: 0,
		},
		{
			name:        "constrained with enum",
			schemaPath:  "testdata/strings_should_be_constrained/correct_enum.json",
			nViolations: 0,
		},
		{
			name:        "constrained with format",
			schemaPath:  "testdata/strings_should_be_constrained/correct_format.json",
			nViolations: 0,
		},
		{
			name:        "constrained with maxLength",
			schemaPath:  "testdata/strings_should_be_constrained/correct_maxlength.json",
			nViolations: 0,
		},
		{
			name:        "constrained with minLength",
			schemaPath:  "testdata/strings_should_be_constrained/correct_minlength.json",
			nViolations: 0,
		},
		{
			name:        "constrained with pattern",
			schemaPath:  "testdata/strings_should_be_constrained/correct_pattern.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			rule := StringsShouldBeConstrained{}
			ruleResults := rule.Verify(schema)

			if len(ruleResults.Violations) != tc.nViolations {
				violationNames := []string{}
				for _, violation := range ruleResults.Violations {
					violationNames = append(violationNames, violation.Message)
				}
				t.Fatalf(
					"Unexpected number of rule violations in test case '%s': Expected %d, got %d. Got these: %v. Schema: %#v",
					tc.name,
					tc.nViolations,
					len(ruleResults.Violations),
					violationNames,
					schema.Schema,
				)
			}
		})
	}
}
