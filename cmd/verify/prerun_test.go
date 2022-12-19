package verify

import "testing"

func TestVerifyFlags(t *testing.T) {
	testCases := []struct {
		ruleSet              string
		skipNormalization    bool
		skipSchemaValidation bool
		expectedError        bool
	}{
		{
			ruleSet:              "",
			skipNormalization:    false,
			skipSchemaValidation: false,
			expectedError:        false,
		}, {
			ruleSet:              "cluster-app",
			skipNormalization:    false,
			skipSchemaValidation: false,
			expectedError:        false,
		}, {
			ruleSet:              "non-existing",
			skipNormalization:    false,
			skipSchemaValidation: false,
			expectedError:        true,
		}, {
			ruleSet:              "",
			skipNormalization:    true,
			skipSchemaValidation: true,
			expectedError:        true,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			flag := &flag{
				ruleSet:              tc.ruleSet,
				skipNormalization:    tc.skipNormalization,
				skipSchemaValidation: tc.skipSchemaValidation,
			}

			err := flag.verify()
			if err != nil && !tc.expectedError {
				t.Fatalf("Unexpected error: %s", err)
			}
			if err == nil && tc.expectedError {
				t.Fatalf("Expected error, got none")
			}
		})
	}
}
