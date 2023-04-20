package recurse

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/schema"
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
			s, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			goldFoundStruct := &struct {
				found bool
				path  string
			}{
				found: false,
				path:  "",
			}
			checkForGold := func(s *schema.ExtendedSchema) {
				if s.Title == "Gold" {
					goldFoundStruct.found = true
					goldFoundStruct.path = s.GetResolvedLocation()
				}
			}

			RecurseAll(s, checkForGold)
			if !goldFoundStruct.found {
				t.Fatalf("Expected to find property with title 'Gold'")
			}
			if goldFoundStruct.path != tc.goldPath {
				t.Fatalf("Expected path '%s', got '%s'", tc.goldPath, goldFoundStruct.path)
			}
		})
	}
}

func TestSelfReferencingRecurse(t *testing.T) {
	s, err := schema.Compile("testdata/self_referencing_ref.json")
	if err != nil {
		t.Fatalf("Unexpected parsing error: %s", err)
	}

	dummyFunc := func(s *schema.ExtendedSchema) {}
	// if this does not loop forever, the test passes
	RecurseAll(s, dummyFunc)
}
