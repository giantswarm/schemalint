package rules

import (
	"testing"

	"github.com/giantswarm/schemalint/pkg/lint"
)

func TestDescriptionCorrect(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
		rules       []lint.Rule
	}{
		{
			name:        "description contains line breaks",
			schemaPath:  "testdata/description_correct/description_with_illegal_chars.json",
			nViolations: 6,
			rules:       []lint.Rule{DescriptionMustNotContainIllegalCharacters{}},
		},
		{
			name:        "description is not sentence case",
			schemaPath:  "testdata/description_correct/description_not_sentence_case.json",
			nViolations: 1,
			rules:       []lint.Rule{DescriptionMustBeSentenceCase{}},
		},
		{
			name:        "description contains title",
			schemaPath:  "testdata/description_correct/description_contains_title.json",
			nViolations: 1,
			rules:       []lint.Rule{DescriptionShouldNotContainTitle{}},
		},
		{
			name:        "description is too short",
			schemaPath:  "testdata/description_correct/description_too_short.json",
			nViolations: 1,
			rules:       []lint.Rule{DescriptionShouldHaveCorrectLength{}},
		},
		{
			name:        "description is too long",
			schemaPath:  "testdata/description_correct/description_too_long.json",
			nViolations: 1,
			rules:       []lint.Rule{DescriptionShouldHaveCorrectLength{}},
		},
		{
			name:        "description does not use punctuation",
			schemaPath:  "testdata/description_correct/description_no_punctuation.json",
			nViolations: 1,
			rules:       []lint.Rule{DescriptionMustUsePunctuation{}},
		},
		{
			name:        "description is missing",
			schemaPath:  "testdata/description_correct/8_missing_descriptions.json",
			nViolations: 0,
			rules: []lint.Rule{
				DescriptionMustNotContainIllegalCharacters{},
				DescriptionMustBeSentenceCase{},
				DescriptionShouldNotContainTitle{},
				DescriptionShouldHaveCorrectLength{},
				DescriptionMustUsePunctuation{},
			},
		},
		{
			name:        "description is correct",
			schemaPath:  "testdata/description_correct/description_correct.json",
			nViolations: 0,
			rules: []lint.Rule{
				DescriptionMustNotContainIllegalCharacters{},
				DescriptionMustBeSentenceCase{},
				DescriptionShouldNotContainTitle{},
				DescriptionShouldHaveCorrectLength{},
				DescriptionMustUsePunctuation{},
			},
		},
		{
			name:        "all rules fail",
			schemaPath:  "testdata/description_correct/description_all_rules_fail.json",
			nViolations: 5,
			rules: []lint.Rule{
				DescriptionMustNotContainIllegalCharacters{},
				DescriptionMustBeSentenceCase{},
				DescriptionShouldNotContainTitle{},
				DescriptionShouldHaveCorrectLength{},
				DescriptionMustUsePunctuation{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := lint.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			ruleResults := []string{}
			for _, rule := range tc.rules {
				ruleResults = append(ruleResults, rule.Verify(schema).Violations...)
			}

			if len(ruleResults) != tc.nViolations {
				t.Fatalf("Unexpected number of rule violations in test case '%s': Expected %d, got %d", tc.name, tc.nViolations, len(ruleResults))
			}
		})
	}
}
