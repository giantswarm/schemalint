package utils

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type RecurseType int

const (
	TypeRecurseAll RecurseType = iota
	TypeRecurseProperties
	TypeRecursePropertiesWithDescription
)

func TestRecurse(t *testing.T) {
	testCases := []struct {
		name       string
		schemaPath string
		goldPath   string
	}{
		{
			name:       "$ref to external reference",
			schemaPath: "testdata/with_external_ref.json",
			goldPath:   "/properties/address/properties/gold",
		},
		{
			name:       "$ref to internal entry in $defs",
			schemaPath: "testdata/with_internal_defs_ref.json",
			goldPath:   "/properties/refProperty/properties/gold",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			getGoldPath := func(schema *schemautils.ExtendedSchema) []string {
				if schema.Title == "Gold" {
					return []string{schema.GetResolvedLocation()}
				}
				return []string{}
			}

			paths := RecurseAll(schema, getGoldPath)
			if len(paths) != 1 {
				t.Fatalf("Expected 1 path, got %d", len(paths))
			}
			if paths[0] != tc.goldPath {
				t.Fatalf("Expected path '%s', got '%s'", tc.goldPath, paths[0])
			}
		})
	}
}

func TestSelfReferencingRecurse(t *testing.T) {
	schema, err := lint.Compile("testdata/self_referencing_ref.json")
	if err != nil {
		t.Fatalf("Unexpected parsing error: %s", err)
	}

	dummyFunc := func(schema *schemautils.ExtendedSchema) []string {
		return []string{}
	}
	// if this does not loop forever, the test passes
	RecurseAll(schema, dummyFunc)
}
