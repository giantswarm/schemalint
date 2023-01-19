package normalize

import "testing"

func TestValidateFlags(t *testing.T) {
	testCases := []struct {
		name           string
		outputPath     string
		forceOverwrite bool
		expectedError  bool
	}{
		{
			name:           "no output path",
			outputPath:     "",
			forceOverwrite: false,
			expectedError:  false,
		}, {
			name:           "output path - overwrite - no file",
			outputPath:     "testdata/does-not-exist",
			forceOverwrite: true,
			expectedError:  false,
		}, {
			name:           "output path - overwrite - file exists",
			outputPath:     "testdata/does-not-exist",
			forceOverwrite: true,
			expectedError:  false,
		}, {
			name:           "output path - no overwrite - file exists",
			outputPath:     "testdata/file",
			forceOverwrite: false,
			expectedError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			flag := &flag{
				outputPath:     tc.outputPath,
				forceOverwrite: tc.forceOverwrite,
			}
			err := flag.validate()
			if err != nil && !tc.expectedError {
				t.Fatalf("Unexpected error in test case '%s': %s", tc.name, err)
			}
			if err == nil && tc.expectedError {
				t.Fatalf("Expected error, got none in test case '%s'", tc.name)
			}
		})
	}
}
