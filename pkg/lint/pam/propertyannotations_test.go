package pam

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

func TestBuildPropertyAnnotationsMap(t *testing.T) {
	testCases := []struct {
		name       string
		schemaPath string
		goldenPath string
	}{
		{
			name:       "one property - flat",
			schemaPath: "testdata/flat_simple.json",
			goldenPath: "testdata/flat_simple.golden.json",
		},
		{
			name:       "multiple nested properties",
			schemaPath: "testdata/nested.json",
			goldenPath: "testdata/nested.golden.json",
		},
		{
			name:       "multiple nested properties through reference",
			schemaPath: "testdata/reference_nested.json",
			goldenPath: "testdata/reference_nested.golden.json",
		},
		{
			name:       "multiple nested properties through reference with overriden title",
			schemaPath: "testdata/reference_nested_override.json",
			goldenPath: "testdata/reference_nested_override.golden.json",
		},
		{
			name:       "multiple nested properties through reference with overriden title",
			schemaPath: "testdata/depth_3_simple.json",
			goldenPath: "testdata/depth_3_simple.golden.json",
		},
		{
			name:       "multiple nested properties through reference with overriden title",
			schemaPath: "testdata/depth_equal_prio.json",
			goldenPath: "testdata/depth_equal_prio.golden.json",
		},
		{
			name:       "multiple nested properties through reference with overriden title",
			schemaPath: "testdata/depth_3_root_empty.json",
			goldenPath: "testdata/depth_3_root_empty.golden.json",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			propertyAnnotationsMap := BuildPropertyAnnotationsMap(schema)

			expectedPropertyAnnotationsMap, err := parseGoldenFile(tc.goldenPath)
			if err != nil {
				t.Fatalf("Unexpected error parsing golden file in test case '%s': %s", tc.name, err)
			}

			if !cmp.Equal(propertyAnnotationsMap, expectedPropertyAnnotationsMap) {
				t.Fatalf(
					"Unexpected property annotations map: %s",
					cmp.Diff(propertyAnnotationsMap, expectedPropertyAnnotationsMap),
				)
			}
		})
	}
}

func parseGoldenFile(path string) (PropertyAnnotationsMap, error) {
	goldenFile, err := openGoldenFile(path)
	if err != nil {
		return nil, err
	}

	var propertyAnnotationsMap PropertyAnnotationsMap
	err = json.NewDecoder(goldenFile).Decode(&propertyAnnotationsMap)
	if err != nil {
		return nil, err
	}

	return propertyAnnotationsMap, nil
}

func openGoldenFile(path string) (*os.File, error) {
	path = filepath.Clean(path)
	goldenFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return goldenFile, nil
}

// use this function as a template when creating golden files for new tests
func DumpPropertyAnnotationsMap(propertyAnnotationsMap PropertyAnnotationsMap, path string) {
	path = filepath.Clean(path)
	goldenFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	json, err := json.MarshalIndent(propertyAnnotationsMap, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = goldenFile.Write(json)
	if err != nil {
		panic(err)
	}
}
