package rulesets

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint/rules"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type RuleSet struct {
	rules            []rules.Rule
	excludeLocations []string
	referenceURL     string
}
type RuleSetName string

const (
	NameClusterApp RuleSetName = "cluster-app"
)

var ClusterApp = &RuleSet{
	rules: []rules.Rule{
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
		rules.AdheresToCommonSchemaStructureRecommendations{},
		rules.AdheresToCommonSchemaStructureRequirements{},
		rules.ArrayItemsMustHaveSingleType{},
		rules.AvoidRecursion{},
		rules.AvoidRecursionKeywords{},
		rules.ObjectsMustHaveProperties{},
	},
	excludeLocations: []string{
		"/properties/global/properties/apps",
		"/properties/internal",
		"/properties/providerIntegration",
		"/properties/cluster",
		"/properties/cluster-shared",
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

func Verify(
	name string,
	s *schema.ExtendedSchema,
) (errors rules.RuleResults, recommendations rules.RuleResults) {
	nameEnum := RuleSetName(name)
	ruleSet := GetRuleSet(nameEnum)

	return LintWithRuleSet(s, ruleSet)
}

func GetRuleSetReferenceURL(name string) string {
	nameEnum := RuleSetName(name)
	ruleSet := GetRuleSet(nameEnum)

	return ruleSet.referenceURL
}

func LintWithRuleSet(
	s *schema.ExtendedSchema,
	ruleSet *RuleSet,
) (errors rules.RuleResults, recommendations rules.RuleResults) {
	errors = rules.RuleResults{}
	recommendations = rules.RuleResults{}
	for _, rule := range ruleSet.rules {
		ruleResults := rule.Verify(s)
		severity := rule.GetSeverity()

		filteredResults := ruleResults.Filter(ruleSet.excludeLocations)

		switch severity {
		case rules.SeverityError:
			errors.Concat(filteredResults)
		case rules.SeverityRecommendation:
			recommendations.Concat(filteredResults)
		}
	}

	return errors, recommendations
}
