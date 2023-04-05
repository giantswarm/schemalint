package rulesets

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
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
			nRecommendations: 146,
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
			nRecommendations: 11,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			errors, recommendations := VerifyRuleSet(tc.ruleSetName, schema)
			if len(errors) != tc.nErrors {
				t.Fatalf(
					"Unexpected number of errors in test case '%s': Expected %d, got %d",
					tc.name,
					tc.nErrors,
					len(errors),
				)
			}
			if len(recommendations) != tc.nRecommendations {
				t.Fatalf(
					"Unexpected number of recommendations in test case '%s': Expected %d, got %d",
					tc.name,
					tc.nRecommendations,
					len(recommendations),
				)
			}
		})
	}
}
