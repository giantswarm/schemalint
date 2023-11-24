package rulesets

import (
	"fmt"
	"strings"
	"testing"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

func TestLintWithRules(t *testing.T) {
	testCases := []struct {
		name             string
		schemaPath       string
		ruleSetName      string
		nErrors          int
		nRecommendations int
	}{
		{
			name:             "case 0: cluster-app - cluster-azure",
			schemaPath:       "testdata/cluster_azure.json",
			ruleSetName:      "cluster-app",
			nErrors:          0,
			nRecommendations: 92,
		},
		{
			name:             "case 1: with ignored locations",
			schemaPath:       "testdata/with_ignored.json",
			ruleSetName:      "cluster-app",
			nErrors:          4,
			nRecommendations: 2,
		},
		{
			name:             "case 2: without ignored locations",
			schemaPath:       "testdata/no_ignored.json",
			ruleSetName:      "cluster-app",
			nErrors:          9,
			nRecommendations: 8,
		},
		{
			name:             "case 3: all rules violated",
			schemaPath:       "testdata/all_rules_violated.json",
			ruleSetName:      "cluster-app",
			nErrors:          42,
			nRecommendations: 30,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			errors, recommendations := Verify(tc.ruleSetName, s)
			if len(errors.Violations) != tc.nErrors {
				var errorMessages []string
				for i, violation := range errors.Violations {
					errorMessages = append(
						errorMessages,
						fmt.Sprintf("%d. %s: %s", i+1, violation.Location, violation.Message))
				}
				mergedMessages := strings.Join(errorMessages, "\n")

				t.Fatalf(
					"Unexpected number of errors in test case '%s': Expected %d, got %d\n\nErrors:\n%s",
					tc.name,
					tc.nErrors,
					len(errors.Violations),
					mergedMessages,
				)
			}
			if len(recommendations.Violations) != tc.nRecommendations {
				var recommendationMessages []string
				for i, recommendation := range recommendations.Violations {
					recommendationMessages = append(
						recommendationMessages,
						fmt.Sprintf("%d. %s: %s", i+1, recommendation.Location, recommendation.Message))
				}
				mergedMessages := strings.Join(recommendationMessages, "\n")

				t.Fatalf(
					"Unexpected number of recommendations in test case '%s': Expected %d, got %d\n\nRecommendations:\n%s",
					tc.name,
					tc.nRecommendations,
					len(recommendations.Violations),
					mergedMessages,
				)
			}
		})
	}
}
