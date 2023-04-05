package normalize

import "testing"

func TestIsUnchanged(t *testing.T) {
	testCases := []struct {
		name        string
		output      []byte
		outputPath  string
		isUnchanged bool
	}{
		{
			name:        "case 0: unchanged",
			output:      []byte("test\n"),
			outputPath:  "testdata/is_unchanged/unchanged",
			isUnchanged: true,
		}, {
			name:       "case 1: changed",
			output:     []byte("test\n"),
			outputPath: "testdata/is_unchanged/changed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isUnchanged := checkIsUnchanged(tc.output, tc.outputPath)
			if isUnchanged != tc.isUnchanged {
				t.Fatalf("Unexpected result in test case '%s': Expected %t, got %t", tc.name, tc.isUnchanged, isUnchanged)
			}
		})
	}
}
