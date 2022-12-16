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
		{
			// If this test happens to fail, it might be that our linter
			// actually changed and does support negative lookbehind.
			// See https://github.com/santhosh-tekuri/jsonschema/pull/60
			name:      "Regex",
			inputPath: "testdata/regex-negative-lookbehind.json",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Compile(tt.inputPath); (err != nil) != tt.wantErr {
				t.Errorf("CompileAuto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
