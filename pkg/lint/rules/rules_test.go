package rules

import "testing"

func TestFilter(t *testing.T) {
	testCases := []struct {
		name              string
		ruleResults       *RuleResults
		excludedLocations []string
		nViolations       int
	}{
		{
			name: "no exlusions",
			ruleResults: &RuleResults{
				Violations: []Violation{
					{
						Location: "properties/noInternal",
						Message:  "Title missing",
					},
				},
			},
			excludedLocations: []string{"properties/internal"},
			nViolations:       1,
		},
		{
			name: "exclusions",
			ruleResults: &RuleResults{
				Violations: []Violation{
					{
						Location: "properties/internal",
						Message:  "Title missing",
					},
				},
			},
			excludedLocations: []string{"properties/internal"},
			nViolations:       0,
		},
		{
			name: "no exlusions, but property name is a substring of excluded location",
			ruleResults: &RuleResults{
				Violations: []Violation{
					{
						Location: "properties/internal2",
						Message:  "Title missing",
					},
				},
			},
			excludedLocations: []string{"properties/internal"},
			nViolations:       1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filtered := tc.ruleResults.Filter(tc.excludedLocations)
			if len(filtered.Violations) != tc.nViolations {
				t.Fatalf(
					"Unexpected number of violations: %d, expected %d",
					len(filtered.Violations),
					tc.nViolations,
				)
			}
		})
	}
}
