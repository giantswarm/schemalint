package lint

import (
	"testing"

	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
)

func TestCompileAuto(t *testing.T) {
	tests := []struct {
		name      string
		inputPath string
		wantErr   bool
	}{
		{
			name:      "Minimal valid",
			inputPath: "testdata/valid_2020-12.json",
			wantErr:   false,
		},
		{
			name:      "Minimal invalid",
			inputPath: "testdata/schema-url-404.json",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := ToFileURL(tt.inputPath)
			if err != nil {
				t.Errorf("CompileAuto() error creating file URL: %s", err)
			}
			if err := CompileAuto(url); (err != nil) != tt.wantErr {
				t.Errorf("CompileAuto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
