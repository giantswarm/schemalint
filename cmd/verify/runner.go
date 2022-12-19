package verify

import (
	"errors"
	"fmt"
	"os"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/findings"
	"github.com/giantswarm/schemalint/pkg/lint/rulesets"
	"github.com/giantswarm/schemalint/pkg/normalize"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"
)

type runner struct {
	flag *flag
}

// Structure to collect results from different checks
type TestResult struct {
	Name     string
	Success  bool
	Findings []findings.LintFindings
}

func (r *runner) run(cmd *cobra.Command, args []string) {
	path := args[0]

	results, success := r.runVerificatonSteps(path)

	printOutput(results)

	if !success {
		os.Exit(1)
	}

}

func (r *runner) runVerificatonSteps(path string) ([]TestResult, bool) {
	results := []TestResult{}
	var schema *jsonschema.Schema

	flags := r.flag

	if !flags.skipSchemaValidation {
		var result TestResult
		result, schema = verifySchemaValidity(path)
		results = append(results, result)
	}
	success := isSuccessful(results)
	if !success {
		return results, success
	}

	if !flags.skipNormalization {
		result := verifyNormalization(path)
		results = append(results, result)
	}

	if !flags.skipSchemaValidation && flags.ruleSet != "" {
		result := verifyRuleSet(flags.ruleSet, schema)
		results = append(results, result)

	}
	success = isSuccessful(results)
	return results, success
}

func verifySchemaValidity(path string) (TestResult, *jsonschema.Schema) {
	schema, err := lint.Compile(path)

	lintFindings := []findings.LintFindings{}
	if err != nil {
		lintFindings = append(lintFindings, findings.LintFindings{
			Message:  err.Error(),
			Severity: findings.SeverityError,
		})
	}

	result := TestResult{
		Name:     "Is valid JSON schema",
		Success:  err == nil,
		Findings: lintFindings,
	}

	return result, schema
}

func verifyNormalization(path string) TestResult {
	var err error
	var content []byte

	content, err = os.ReadFile(path)
	if err == nil {
		var isNormalized bool
		isNormalized, err = normalize.CheckIsNormalized(content)
		if err == nil && !isNormalized {
			err = errors.New("schema is not normalized")
		}

	}
	lintFindings := []findings.LintFindings{}
	if err != nil {
		lintFindings = append(lintFindings, findings.LintFindings{
			Message:  err.Error(),
			Severity: findings.SeverityError,
		})
	}

	result := TestResult{
		Name:     "Is normalized",
		Success:  err == nil,
		Findings: lintFindings,
	}
	return result

}

func verifyRuleSet(ruleSet string, schema *jsonschema.Schema) TestResult {
	lintFindings := rulesets.VerifyRuleSet(ruleSet, schema)

	result := TestResult{
		Name:     fmt.Sprintf("Rule set '%s'", ruleSet),
		Success:  findings.GetCount(lintFindings, findings.SeverityError) == 0,
		Findings: lintFindings,
	}

	return result
}

func isSuccessful(results []TestResult) bool {
	for _, result := range results {
		if !result.Success {
			return false
		}
	}
	return true
}
