package schemautils_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

func TestGetSchemasAt(t *testing.T) {
	testCases := []struct {
		name           string
		schemaPath     string
		location       string
		expectedTitles []string
	}{
		{
			name:           "simple",
			schemaPath:     "testdata/simple.json",
			location:       "/properties/rootProp",
			expectedTitles: []string{"gold"},
		},
		{
			name:           "nested",
			schemaPath:     "testdata/nested.json",
			location:       "/properties/rootProp/properties/childProp/properties/grandchildProp",
			expectedTitles: []string{"gold"},
		},
		{
			name:           "referenced",
			schemaPath:     "testdata/referenced.json",
			location:       "/properties/rootProp/properties/childProp/properties/grandchildProp",
			expectedTitles: []string{"gold", "gold_ref", "gold_ref_ref"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}
			foundSchemas := schema.GetSchemasAt(tc.location)
			foundTitles := []string{}
			for _, foundSchema := range foundSchemas {
				foundTitles = append(foundTitles, foundSchema.Title)
			}
			if !cmp.Equal(foundTitles, tc.expectedTitles) {
				t.Fatalf("Unexpected schemas found in test case '%s':\n%s", tc.name, cmp.Diff(foundTitles, tc.expectedTitles))
			}
		})
	}
}

func TestGetParentPropertyPath(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "case 0: two level",
			input:         "/properties/rootProp/properties/childProp",
			expected:      "/properties/rootProp",
			expectedError: false,
		},
		{
			name:          "case 1: three level",
			input:         "/properties/rootProp/properties/childProp/properties/grandchildProp",
			expected:      "/properties/rootProp/properties/childProp",
			expectedError: false,
		},
		{
			name:          "case 2: root",
			input:         "/properties/rootProp",
			expected:      "",
			expectedError: false,
		},
		{
			name:          "case 3: no property",
			input:         "/properties",
			expected:      "",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := schemautils.GetParentPropertyPath(tc.input)
			if tc.expectedError && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tc.expectedError && err != nil {
				t.Fatalf("expected no error but got one: %s", err)
			}

			if actual != tc.expected {
				t.Fatalf("Unexpected parent path in test case '%s':\n%s", tc.name, cmp.Diff(actual, tc.expected))
			}
		})
	}
}
