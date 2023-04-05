package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestObjectsMustHaveProperties(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
	}{
		{
			name:        "case 0: object without properties",
			schemaPath:  "testdata/objects_must_have_properties/object_without_properties.json",
			nViolations: 1,
		},
		{
			name:        "case 1: object with property",
			schemaPath:  "testdata/objects_must_have_properties/object_with_property.json",
			nViolations: 0,
		},
		{
			name:        "case 2: object with additional property",
			schemaPath:  "testdata/objects_must_have_properties/object_with_additional_property.json",
			nViolations: 0,
		},
		{
			name:        "case 3: object with pattern property",
			schemaPath:  "testdata/objects_must_have_properties/object_with_pattern_property.json",
			nViolations: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			rule := ObjectsMustHaveProperties{}
			ruleResults := rule.Verify(schema).Violations

			if len(ruleResults) != tc.nViolations {
				t.Fatalf(
					"Unexpected number of rule violations in test case '%s': Expected %d, got %d",
					tc.name,
					tc.nViolations,
					len(ruleResults),
				)
			}
		})
	}
}
