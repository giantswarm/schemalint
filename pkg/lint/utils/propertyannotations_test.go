package utils

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/giantswarm/schemalint/pkg/lint"
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			propertyAnnotationsMap := BuildPropertyAnnotationsMap(schema)

			expectedPropertyAnnotationsMap, err := parseGoldenFile(tc.goldenPath)
			if err != nil {
				t.Fatalf("Unexpected error parsing golden file in test case '%s': %s", tc.name, err)
			}
			// dumpPropertyAnnotationsMap(propertyAnnotationsMap)

			if !cmp.Equal(propertyAnnotationsMap, expectedPropertyAnnotationsMap) {
				t.Fatalf("Unexpected property annotations map: %s", cmp.Diff(propertyAnnotationsMap, expectedPropertyAnnotationsMap))
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
	goldenFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return goldenFile, nil
}

func dumpPropertyAnnotationsMap(propertyAnnotationsMap PropertyAnnotationsMap) {
	json, err := json.MarshalIndent(propertyAnnotationsMap, "", "  ")
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(json)
}
