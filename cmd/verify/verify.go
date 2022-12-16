package verify

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/lint/findings"
)

var (
	VerifyCmd = &cobra.Command{
		Use:          "verify PATH",
		Short:        "Verify the given JSON schema input",
		Args:         cobra.ExactArgs(1),
		ArgAliases:   []string{"PATH"},
		PreRun:       preRun,
		Run:          run,
		SilenceUsage: true,
	}

	skipNormalization    bool
	skipSchemaValidation bool

	ruleSet string
)

func init() {
	VerifyCmd.Flags().BoolVar(&skipNormalization, "skip-normalization", false, "Disable the normalization check.")
	VerifyCmd.Flags().BoolVar(&skipSchemaValidation, "skip-schema-validation", false, "Disable the JSON schema validation.")
	VerifyCmd.Flags().StringVar(&ruleSet, "rule-set", "", "The rule set to use for validation.")
}

// Structure to collect results from different checks
type TestResult struct {
	Name     string
	Success  bool
	Findings []findings.LintFindings
}

func run(cmd *cobra.Command, args []string) {
	path := args[0]

	results, success := runVerificatonSteps(path)

	printOutput(results)

	if !success {
		os.Exit(1)
	}
}
