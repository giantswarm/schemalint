package verify

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint/rulesets"
)

func TestValidateFlags(t *testing.T) {
	testCases := []struct {
		name                 string
		skipNormalization    bool
		skipSchemaValidation bool
		ruleSet              string
		expectedError        bool
	}{
		{
			name:                 "no skip - no ruleset",
			ruleSet:              "",
			skipNormalization:    false,
			skipSchemaValidation: false,
			expectedError:        false,
		},
		{
			name:                 "skip both",
			ruleSet:              "",
			skipNormalization:    true,
			skipSchemaValidation: true,
			expectedError:        true,
		},
		{
			name:                 "existing ruleset",
			ruleSet:              string(rulesets.NameClusterApp),
			skipNormalization:    false,
			skipSchemaValidation: false,
			expectedError:        false,
		},
		{
			name:                 "non-existing ruleset",
			ruleSet:              "non-existing",
			skipNormalization:    false,
			skipSchemaValidation: false,
			expectedError:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			flag := &flag{
				ruleSet:              tc.ruleSet,
				skipNormalization:    tc.skipNormalization,
				skipSchemaValidation: tc.skipSchemaValidation,
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
