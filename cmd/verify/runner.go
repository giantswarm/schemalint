package verify

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/schema"
	"github.com/giantswarm/schemalint/pkg/verify"
)

type runner struct {
	flag *flag
}

func (r *runner) run(cmd *cobra.Command, args []string) {
	path := args[0]

	results, success := r.runVerificatonSteps(path)

	printOutput(results)

	if !success {
		os.Exit(1)
	}

}

func (r *runner) runVerificatonSteps(path string) ([]verify.TestResult, bool) {
	results := []verify.TestResult{}
	var s *schema.ExtendedSchema

	flags := r.flag

	if !flags.skipSchemaValidation {
		var result verify.TestResult
		result, s = verify.CheckSchemaValidity(path)
		results = append(results, result)
	}
	success := isSuccessful(results)
	if !success {
		return results, success
	}

	if !flags.skipNormalization {
		result := verify.CheckNormalization(path)
		results = append(results, result)
	}

	if !flags.skipSchemaValidation && flags.ruleSet != "" {
		result := verify.CheckRuleSet(flags.ruleSet, s)
		results = append(results, result)

	}
	success = isSuccessful(results)
	return results, success
}

func isSuccessful(results []verify.TestResult) bool {
	for _, result := range results {
		if !result.Success {
			return false
		}
	}
	return true
}
