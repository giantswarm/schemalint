package verify

import (
	"errors"

	"github.com/giantswarm/schemalint/pkg/cli"
	"github.com/giantswarm/schemalint/pkg/lint/rulesets"
	"github.com/spf13/cobra"
)

type flag struct {
	skipNormalization    bool
	skipSchemaValidation bool
	ruleSet              string
}

func (f *flag) init(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&f.skipNormalization, "skip-normalization", false, "Disable the normalization check.")
	cmd.Flags().BoolVar(&f.skipSchemaValidation, "skip-schema-validation", false, "Disable the JSON schema validation.")
	cmd.Flags().StringVar(&f.ruleSet, "rule-set", "", "The rule set to use for validation.")
}

func (f *flag) verify() []error {
	var errors []error

	if err := verifySkipFlags(f); err != nil {
		errors = append(errors, err)
	}
	if err := verifyRuleSetFlag(f); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func verifySkipFlags(flag *flag) error {
	if flag.skipNormalization && flag.skipSchemaValidation {
		return errors.New(cli.SprintErrorMessage("both --skip-normalization and --skip-schema-validation are set, so we have no checks to run"))
	}
	return nil
}

func verifyRuleSetFlag(flag *flag) error {
	if flag.ruleSet != "" && !rulesets.IsRuleSetName(flag.ruleSet) {
		availableRuleSets := rulesets.GetAvailableRuleSets()
		return errors.New(cli.SprintfErrorMessage("unknown rule set %s, available rule sets are: %v", flag.ruleSet, availableRuleSets))
	}
	return nil
}
