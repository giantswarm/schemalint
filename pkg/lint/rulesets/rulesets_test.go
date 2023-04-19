package rulesets

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/schema"
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
			nRecommendations: 1,
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
				t.Fatalf(
					"Unexpected number of errors in test case '%s': Expected %d, got %d",
					tc.name,
					tc.nErrors,
					len(errors.Violations),
				)
			}
			if len(recommendations.Violations) != tc.nRecommendations {
				t.Fatalf(
					"Unexpected number of recommendations in test case '%s': Expected %d, got %d",
					tc.name,
					tc.nRecommendations,
					len(recommendations.Violations),
				)
			}
		})
	}
}
