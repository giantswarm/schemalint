package rulesets

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint/rulesmeta"
	"github.com/giantswarm/schemalint/pkg/lint/rulesmeta/rules"
)

type RuleSet struct {
	rules []rulesmeta.Rule
}
type RuleSetName string

const (
	NameClusterApp RuleSetName = "cluster-app"
)

var ClusterApp = &RuleSet{
	rules: []rulesmeta.Rule{
		rules.TitleExists{},
		rules.DescriptionExists{},
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

	errors = []string{}
	recommendations = []string{}
	for _, rule := range ruleSet.rules {
		violations := rule.Verify(schema)
		severity := rule.GetSeverity()
		if severity == rulesmeta.SeverityError {
			errors = append(errors, violations...)
		}
		if severity == rulesmeta.SeverityRecomendation {
			recommendations = append(recommendations, violations...)
		}
	}

	return errors, recommendations
}
