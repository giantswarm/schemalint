package rules

import (
	"fmt"
	"testing"

	"github.com/giantswarm/schemalint/pkg/schema"
)

func TestTitleCorrect(t *testing.T) {
	testCases := []struct {
		name        string
		schemaPath  string
		nViolations int
		rules       []Rule
	}{
		{
			name:        "title contains line breaks",
			schemaPath:  "testdata/title_correct/title_with_illegal_chars.json",
			nViolations: 10,
			rules:       []Rule{TitleMustNotContainIllegalCharacters{}},
		},
		{
			name:        "title is not sentence case",
			schemaPath:  "testdata/title_correct/title_not_sentence_case.json",
			nViolations: 1,
			rules:       []Rule{TitleMustBeSentenceCase{}},
		},
		{
			name:        "title should not contain the parents title",
			schemaPath:  "testdata/title_correct/title_contains_parents_title.json",
			nViolations: 1,
			rules:       []Rule{TitleShouldNotContainParentsTitle{}},
		},
		{
			name:        "title is correct",
			schemaPath:  "testdata/title_correct/title_correct.json",
			nViolations: 0,
			rules: []Rule{
				TitleMustNotContainIllegalCharacters{},
				TitleMustBeSentenceCase{},
				TitleShouldNotContainParentsTitle{},
			},
		},
		{
			name:        "all rules fail",
			schemaPath:  "testdata/title_correct/title_all_rules_fail.json",
			nViolations: 4,
			rules: []Rule{
				TitleMustNotContainIllegalCharacters{},
				TitleMustBeSentenceCase{},
				TitleShouldNotContainParentsTitle{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := schema.Compile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Unexpected parsing error in test case '%s': %s", tc.name, err)
			}

			ruleResults := []Violation{}
			for _, rule := range tc.rules {
				ruleResults = append(ruleResults, rule.Verify(schema).Violations...)
			}

			if len(ruleResults) != tc.nViolations {
				for _, violation := range ruleResults {
					fmt.Println(violation)
				}
				t.Fatalf(
					"Unexpected number of rule violations in test case '%s': Expected %d, got %d",
					tc.name,
					tc.nViolations,
					len(ruleResults),
				)
			}
		})
	}
}
