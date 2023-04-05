package verify

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/rulesets"
	"github.com/giantswarm/schemalint/pkg/normalize"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type runner struct {
	flag *flag
}

// Structure to collect results from different checks
type TestResult struct {
	Message         string
	Success         bool
	Errors          []string
	Recommendations []string
	MoreInfo        string
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
	var schema *schemautils.ExtendedSchema

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

func verifySchemaValidity(path string) (TestResult, *schemautils.ExtendedSchema) {
	schema, err := lint.Compile(path)

	compileErrors := []string{}
	if err != nil {
		compileErrors = append(compileErrors, err.Error())
	}

	success := err == nil
	message := "Input is valid JSON Schema."
	if !success {
		message = "Input is not valid JSON Schema."
	}
	result := TestResult{
		Message: message,
		Success: success,
		Errors:  compileErrors,
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
			err = errors.New("Schema is not normalized.")
		}

	}
	errors := []string{}
	if err != nil {
		errors = append(errors, err.Error())
	}

	success := err == nil
	message := "Input is normalized."
	if !success {
		message = "Input is not normalized."
	}
	result := TestResult{
		Message: message,
		Success: success,
		Errors:  errors,
	}
	return result

}

func verifyRuleSet(ruleSet string, schema *schemautils.ExtendedSchema) TestResult {
	errors, recommendations := rulesets.VerifyRuleSet(ruleSet, schema)

	success := len(errors) == 0
	message := fmt.Sprintf("Input is valid according to rule set '%s'.", ruleSet)
	if !success {
		message = fmt.Sprintf("Input is not valid according to rule set '%s'.", ruleSet)
	}
	moreInfo := ""
	if len(recommendations)+len(errors) > 0 {
		moreInfo = fmt.Sprintf(
			"For more information regarding the errors and recommendations, please refer: \"%s\".",
			rulesets.GetRuleSetReferenceURL(ruleSet),
		)
	}

	result := TestResult{
		Message:         message,
		Success:         success,
		Errors:          errors,
		Recommendations: recommendations,
		MoreInfo:        moreInfo,
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
