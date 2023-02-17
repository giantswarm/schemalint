package rulesets

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/rules"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type RuleSet struct {
	rules []lint.Rule
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
		rules.StringsShouldBeConstrained{},
		rules.NumbersShouldBeConstrained{},
		rules.ExampleExists{},
		rules.DeprecatedPropertiesShouldHaveComment{},
		rules.ExamplesShouldNotBeTooMany{},
		rules.PropertiesMustHaveOneType{},
	},
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

func VerifyRuleSet(name string, schema *schemautils.ExtendedSchema) (errors []string, recommendations []string) {
	nameEnum := RuleSetName(name)
	ruleSet := GetRuleSet(nameEnum)

	return LintWithRules(schema, ruleSet.rules)
}

func LintWithRules(schema *schemautils.ExtendedSchema, rules []lint.Rule) (errors []string, recommendations []string) {
	errors = []string{}
	recommendations = []string{}
	for _, rule := range rules {
		ruleResults := rule.Verify(schema)
		severity := rule.GetSeverity()
		if severity == lint.SeverityError {
			errors = append(errors, ruleResults.Violations...)
		}
		if severity == lint.SeverityRecommendation {
			recommendations = append(recommendations, ruleResults.Violations...)
		}
	}

	return errors, recommendations
}
