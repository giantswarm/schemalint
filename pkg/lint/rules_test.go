package lint

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
				Violations: []RuleViolation{
					{
						location: "properties/noInternal",
						message:  "Title missing",
					},
				},
			},
			excludedLocations: []string{"properties/internal"},
			nViolations:       1,
		},
		{
			name: "exclusions",
			ruleResults: &RuleResults{
				Violations: []RuleViolation{
					{
						location: "properties/internal",
						message:  "Title missing",
					},
				},
			},
			excludedLocations: []string{"properties/internal"},
			nViolations:       0,
		},
		{
			name: "no exlusions, but property name is a substring of excluded location",
			ruleResults: &RuleResults{
				Violations: []RuleViolation{
					{
						location: "properties/internal2",
						message:  "Title missing",
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
				t.Fatalf("Unexpected number of violations: %d, expected %d", len(filtered.Violations), tc.nViolations)
			}
		})
	}
}
