package rulesets

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/rules"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type RuleSet struct {
	rules            []lint.Rule
	excludeLocations []string
	referenceURL     string
}
type RuleSetName string

const (
	NameClusterApp RuleSetName = "cluster-app"
)

var ClusterApp = &RuleSet{
	rules: []lint.Rule{
		rules.TitleExists{},
		rules.TitleMustBeSentenceCase{},
		rules.TitleMustNotContainIllegalCharacters{},
		rules.TitleShouldNotContainParentsTitle{},
		rules.DescriptionExists{},
		rules.DescriptionMustNotContainIllegalCharacters{},
		rules.DescriptionMustBeSentenceCase{},
		rules.DescriptionMustUsePunctuation{},
		rules.DescriptionShouldNotContainTitle{},
		rules.DescriptionShouldHaveCorrectLength{},
		rules.MustUseCorrectDialect{},
		rules.ShouldDisableAdditionalProperties{},
		rules.ArraysMustHaveItems{},
		rules.AvoidXOf{},
		rules.StringsShouldBeConstrained{},
		rules.NumbersShouldBeConstrained{},
		rules.ExampleExists{},
		rules.DeprecatedPropertiesShouldHaveComment{},
		rules.ExamplesShouldNotBeTooMany{},
		rules.AvoidUnevaluated{},
		rules.PropertiesMustHaveOneType{},
		rules.AvoidLogicalConstruct{},
		rules.ArrayItemsMustHaveSingleType{},
		rules.AvoidRecursion{},
		rules.AvoidRecursionKeywords{},
	},
	excludeLocations: []string{
		"/properties/internal",
	},
	referenceURL: "https://github.com/giantswarm/rfc/pull/55", // should be updated when PR is merged
}

var RuleSets = map[RuleSetName]*RuleSet{
	NameClusterApp: ClusterApp,
}

func GetAvailableRuleSets() []RuleSetName {
	ruleSets := []RuleSetName{}
	for ruleSetName := range RuleSets {
		ruleSets = append(ruleSets, ruleSetName)
	}

	return ruleSets
}

func GetAvailableRuleSetsAsStrings() []string {
	ruleSets := GetAvailableRuleSets()
	stringRuleSets := []string{}
	for _, ruleSet := range ruleSets {
		stringRuleSets = append(stringRuleSets, string(ruleSet))
	}
	return stringRuleSets
}

func IsRuleSetName(name string) bool {
	for _, ruleSet := range GetAvailableRuleSets() {
		if ruleSet == RuleSetName(name) {
			return true
		}
	}

	return false
}

func GetRuleSet(name RuleSetName) *RuleSet {
	ruleSet, ok := RuleSets[name]
	if !ok {
		// This should never happen, as we validate the rule set name before
		panic(fmt.Sprintf("Rule set '%s' not found", name))
	}

	return ruleSet
}

func VerifyRuleSet(
	name string,
	schema *schemautils.ExtendedSchema,
) (errors []string, recommendations []string) {
	nameEnum := RuleSetName(name)
	ruleSet := GetRuleSet(nameEnum)

	return LintWithRuleSet(schema, ruleSet)
}

func GetRuleSetReferenceURL(name string) string {
	nameEnum := RuleSetName(name)
	ruleSet := GetRuleSet(nameEnum)

	return ruleSet.referenceURL
}

func LintWithRuleSet(
	schema *schemautils.ExtendedSchema,
	ruleSet *RuleSet,
) (errors []string, recommendations []string) {
	errors = []string{}
	recommendations = []string{}
	for _, rule := range ruleSet.rules {
		ruleResults := rule.Verify(schema)
		severity := rule.GetSeverity()

		filteredResults := ruleResults.Filter(ruleSet.excludeLocations)

		if severity == lint.SeverityError {
			errors = append(errors, filteredResults.GetMessages()...)
		}
		if severity == lint.SeverityRecommendation {
			recommendations = append(recommendations, filteredResults.GetMessages()...)
		}
	}

	return errors, recommendations
}
