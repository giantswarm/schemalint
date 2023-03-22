package verify

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/cli"
	"github.com/giantswarm/schemalint/pkg/lint/rulesets"
)

type flag struct {
	skipNormalization    bool
	skipSchemaValidation bool
	ruleSet              string
}

func (f *flag) init(cmd *cobra.Command) {
	ruleSets := rulesets.GetAvailableRuleSetsAsStrings()
	cmd.Flags().BoolVar(&f.skipNormalization, "skip-normalization", false, "Disable the normalization check.")
	cmd.Flags().BoolVar(&f.skipSchemaValidation, "skip-schema-validation", false, "Disable the JSON schema validation.")
	cmd.Flags().StringVar(&f.ruleSet, "rule-set", "", "The rule set to use for validation. Available rule sets are: "+strings.Join(ruleSets, ", "))
}

func (f *flag) validate() []error {
	var errors []error

	if err := validateSkipFlags(f); err != nil {
		errors = append(errors, err)
	}
	if err := validateRuleSetFlag(f); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func validateSkipFlags(flag *flag) error {
	if flag.skipNormalization && flag.skipSchemaValidation {
		return errors.New(cli.SprintErrorMessage("both --skip-normalization and --skip-schema-validation are set, so we have no checks to run"))
	}
	return nil
}

func validateRuleSetFlag(flag *flag) error {
	if flag.ruleSet != "" && !rulesets.IsRuleSetName(flag.ruleSet) {
		availableRuleSets := rulesets.GetAvailableRuleSets()
		return errors.New(cli.SprintfErrorMessage("unknown rule set %s, available rule sets are: %v", flag.ruleSet, availableRuleSets))
	}
	return nil
}
