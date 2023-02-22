package schemautils_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestGetSchemasAt(t *testing.T) {
	tests := []struct {
		name           string
		schemaPath     string
		location       string
		expectedTitles []string
	}{
		{
			name:           "simple",
			schemaPath:     "testdata/simple.json",
			location:       "/rootProp",
			expectedTitles: []string{"gold"},
		},
		{
			name:           "nested",
			schemaPath:     "testdata/nested.json",
			location:       "/rootProp/childProp/grandchildProp",
			expectedTitles: []string{"gold"},
		},
		{
			name:           "referenced",
			schemaPath:     "testdata/referenced.json",
			location:       "/rootProp/childProp/grandchildProp",
			expectedTitles: []string{"gold", "gold_ref", "gold_ref_ref"},
		},
	}

	for _, tc := range tests {
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
