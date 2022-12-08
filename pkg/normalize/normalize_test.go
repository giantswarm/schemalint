// Package normalize providesa a function to process a JSON input and return it in
// normalized form.
package normalize

import (
	"flag"
	"io"
	"os"
	"reflect"
	"testing"
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
			name:       "Fourth, with actual schema",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := os.ReadFile("testdata/" + tt.inputPath)
			if err != nil {
				t.Fatalf("Normalize() error when reading file: %v", err)
			}

			got, err := Normalize(input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			want := goldenValue(t, "testdata/"+tt.goldenPath, got, *update)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Normalize() = %v, want %v", string(got), string(want))
			}
		})
	}
}

func goldenValue(t *testing.T, goldenPath string, actual []byte, update bool) []byte {
	t.Helper()

	f, err := os.OpenFile(goldenPath, os.O_RDWR, 0644)
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
