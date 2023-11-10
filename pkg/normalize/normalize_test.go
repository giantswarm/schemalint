// Package normalize providesa a function to process a JSON input and return it in
// normalized form.
package normalize

import (
	"flag"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testdataPath = "testdata/"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name       string
		inputPath  string
		goldenPath string
		wantErr    bool
	}{
		{
			name:       "First",
			inputPath:  "1.json",
			goldenPath: "1.golden",
			wantErr:    false,
		},
		{
			name:       "Second",
			inputPath:  "2.json",
			goldenPath: "2.golden",
			wantErr:    false,
		},
		{
			name:       "Third",
			inputPath:  "3.json",
			goldenPath: "3.golden",
			wantErr:    false,
		},
		{
			name:       "Fourth, with actual s",
			inputPath:  "4.json",
			goldenPath: "4.golden",
			wantErr:    false,
		},
		{
			name:       "Fifth, trailing newline missing",
			inputPath:  "5.json",
			goldenPath: "5.golden",
			wantErr:    false,
		},
		{
			name:      "Sixth, not JSON",
			inputPath: "6.txt",
			wantErr:   true,
		},
		// Property is defined multiple times. Last value wins.
		{
			name:       "Seventh, duplicate key",
			inputPath:  "7.json",
			goldenPath: "7.golden",
			wantErr:    false,
		},
		{
			name:       "Eighth, don't escape < and >",
			inputPath:  "8.json",
			goldenPath: "8.golden",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := os.ReadFile(testdataPath + tt.inputPath)
			if err != nil {
				t.Fatalf("Normalize() error when reading file: %v", err)
			}

			got, err := Normalize(input)
			var want []byte

			if !tt.wantErr {
				if err != nil {
					t.Errorf("Normalize() unexpected error = %v", err)
					return
				}

				if tt.goldenPath != "" {
					want = goldenValue(t, testdataPath+tt.goldenPath, got, *update)
					if !reflect.DeepEqual(got, want) {
						t.Errorf(
							"Unexpected output from Normalize(): %s",
							cmp.Diff(string(got), string(want)),
						)
					}
				}
			} else {
				// Error expected
				if err == nil {
					t.Error("Normalize() expected error, got nil")
				}
			}
		})
	}
}

func goldenValue(t *testing.T, goldenPath string, actual []byte, update bool) []byte {
	t.Helper()

	f, err := os.OpenFile(goldenPath, os.O_RDWR, 0644)
	if err != nil {
		t.Fatalf("Error opening to file %s: %s", goldenPath, err)
	}
	defer f.Close()

	if update {
		_, err := f.Write(actual)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	content, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	return content
}

func TestCheckIsNormalized(t *testing.T) {
	tests := []struct {
		name         string
		inputPath    string
		isNormalized bool
	}{
		{
			name:         "not normalized",
			inputPath:    "1.json",
			isNormalized: false,
		}, {
			name:         "normalized",
			inputPath:    "1.golden",
			isNormalized: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := os.ReadFile(testdataPath + tt.inputPath)
			if err != nil {
				t.Fatalf("IsNormalized() error when reading file: %v", err)
			}

			isNormalized, error := Verify(input)
			if error != nil {
				t.Errorf("IsNormalized() unexpected error = %v", error)
				return
			}
			if isNormalized != tt.isNormalized {
				t.Errorf("IsNormalized() = %v, want %v", isNormalized, tt.isNormalized)
			}
		})
	}

}
