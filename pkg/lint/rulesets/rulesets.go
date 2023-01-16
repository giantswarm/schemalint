package rulesets

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/rules"
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
		rules.DescriptionExists{},
		rules.DescriptionMustNotContainLineBreaks{},
		rules.DescriptionMustBeSentenceCase{},
		rules.DescriptionMustUsePunctuation{},
		rules.DescriptionShouldNotContainTitle{},
		rules.DescriptionShouldHaveCorrectLength{},
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

func VerifyRuleSet(name string, schema *jsonschema.Schema) (errors []string, recommendations []string) {
	nameEnum := RuleSetName(name)
	ruleSet := GetRuleSet(nameEnum)

	return LintWithRules(schema, ruleSet.rules)
}

func LintWithRules(schema *jsonschema.Schema, rules []lint.Rule) (errors []string, recommendations []string) {
	errors = []string{}
	recommendations = []string{}
	for _, rule := range rules {
		violations := rule.Verify(schema)
		severity := rule.GetSeverity()
		if severity == lint.SeverityError {
			errors = append(errors, violations...)
		}
		if severity == lint.SeverityRecomendation {
			recommendations = append(recommendations, violations...)
		}
	}

	return errors, recommendations
}
