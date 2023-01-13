package rulesets

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint/findings"
	"github.com/giantswarm/schemalint/pkg/lint/rulesets/rules"
)

type RuleSet struct {
	rules []rules.Rule
}
type RuleSetName string

const (
	NameClusterApp RuleSetName = "cluster-app"
)

var ClusterApp = &RuleSet{
	rules: []rules.Rule{
		rules.TitleExists{},
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

func VerifyRuleSet(name string, schema *jsonschema.Schema) []findings.LintFindings {
	nameEnum := RuleSetName(name)
	ruleSet := GetRuleSet(nameEnum)

	lintFindings := []findings.LintFindings{}
	for _, rule := range ruleSet.rules {
		violations := rule.Verify(schema)
		lintFindings = append(
			lintFindings,
			tranformRuleViolationsToFindings(violations, rule.GetSeverity())...,
		)
	}

	return lintFindings
}

func tranformRuleViolationsToFindings(violations []rules.RuleViolation, severity findings.Severity) []findings.LintFindings {
	lintFindings := []findings.LintFindings{}
	for _, violation := range violations {
		lintFindings = append(lintFindings, findings.LintFindings{
			Severity: severity,
			Message:  violation.Reason,
		})
	}

	return lintFindings
}
