package verify

import (
	"fmt"
	"os"

	"github.com/giantswarm/schemalint/pkg/cli"
	"github.com/giantswarm/schemalint/pkg/lint/rulesets"
	"github.com/spf13/cobra"
)

func preRun(cmd *cobra.Command, args []string) {
	errors := verifyFlags()
	if len(errors) > 0 {
		for _, e := range errors {
			cli.PrintErrorMessage(e.Error())
		}
		os.Exit(1)
	}
}

func verifyFlags() []error {
	var errors []error

	if err := verifySkipFlags(); err != nil {
		errors = append(errors, err)
	}
	if err := verifyRuleSetFlag(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func verifySkipFlags() error {
	if skipNormalization && skipSchemaValidation {
		return fmt.Errorf("both --skip-normalization and --skip-schema-validation are set, so we have no checks to run")
	}
	return nil
}

func verifyRuleSetFlag() error {
	if ruleSet != "" && !rulesets.IsRuleSetName(ruleSet) {
		availableRuleSets := rulesets.GetAvailableRuleSets()
		return fmt.Errorf("unknown rule set %s, available rule sets are: %v", ruleSet, availableRuleSets)
	}
	return nil
}
