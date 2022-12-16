package rulesets

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint/findings"
	"github.com/giantswarm/schemalint/pkg/lint/rulesets/rules"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type RuleSet struct {
	rules []rules.Rule
}
type RuleSetName string

const (
	ClusterApp RuleSetName = "cluster-app"
)

func GetAvailableRuleSets() []RuleSetName {
	return []RuleSetName{
		ClusterApp,
	}
}

func IsRuleSetName(name string) bool {
	for _, ruleSet := range GetAvailableRuleSets() {
		if ruleSet == RuleSetName(name) {
			return true
		}
	}

	return false
}

func GetRuleSet(name RuleSetName) (*RuleSet, error) {
	switch name {
	case ClusterApp:
		return &RuleSet{
			rules: []rules.Rule{
				rules.TitleExists{},
			},
		}, nil
	default:
		return nil, fmt.Errorf("unknown ruleset: %s", name)
	}
}

func VerifyRuleSet(name string, schema *jsonschema.Schema) []findings.LintFindings {
	nameEnum := RuleSetName(name)
	ruleSet, err := GetRuleSet(nameEnum)
	if err != nil {
		panic(err)
	}

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
