package verify

import (
	"fmt"
	"os"

	"github.com/giantswarm/schemalint/pkg/lint/rules"
	"github.com/giantswarm/schemalint/pkg/lint/rulesets"
	"github.com/giantswarm/schemalint/pkg/normalize"
	"github.com/giantswarm/schemalint/pkg/schema"
)

type TestResult struct {
	Message         string
	Success         bool
	Errors          []Feedback
	Recommendations []Feedback
	MoreInfo        string
}

type Feedback struct {
	Message  string
	Location string // set to "" if feedback is not location-specific
}

func GroupByMessage(feedback []Feedback) map[string][]Feedback {
	grouped := make(map[string][]Feedback, len(feedback))

	for _, f := range feedback {
		grouped[f.Message] = append(grouped[f.Message], f)
	}

	return grouped
}

func FeedbackFromRuleResults(ruleResults rules.RuleResults) []Feedback {
	feedback := make([]Feedback, 0, len(ruleResults.Violations))
	for _, violation := range ruleResults.Violations {
		feedback = append(
			feedback,
			Feedback{Message: violation.Message, Location: violation.Location},
		)
	}
	return feedback
}

func CheckSchemaValidity(path string) (TestResult, *schema.ExtendedSchema) {
	s, err := schema.Compile(path)

	compileErrors := []Feedback{}
	if err != nil {
		compileErrors = append(compileErrors, Feedback{Message: err.Error()})
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

	return result, s
}

func CheckNormalization(path string) TestResult {
	var err error
	var content []byte

	content, err = os.ReadFile(path)
	if err == nil {
		var isNormalized bool
		isNormalized, err = normalize.Verify(content)
		if err == nil && !isNormalized {
			err = fmt.Errorf(
				"Schema is not normalized. Run 'schemalint normalize %[1]s -o %[1]s --force'.",
				path,
			)
		}
	}
	errors := []Feedback{}
	if err != nil {
		errors = append(errors, Feedback{Message: err.Error()})
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

func CheckRuleSet(ruleSet string, s *schema.ExtendedSchema) TestResult {
	errors, recommendations := rulesets.Verify(ruleSet, s)

	errorFeedback := FeedbackFromRuleResults(errors)
	recommendationFeedback := FeedbackFromRuleResults(recommendations)

	success := errors.IsEmpty()
	message := fmt.Sprintf("Input is valid according to rule set '%s'.", ruleSet)
	if !success {
		message = fmt.Sprintf("Input is not valid according to rule set '%s'.", ruleSet)
	}
	moreInfo := ""
	referenceURL := rulesets.GetRuleSetReferenceURL(ruleSet)
	if len(errorFeedback)+len(recommendationFeedback) > 0 && referenceURL != "" {
		moreInfo = fmt.Sprintf(
			"For more information regarding the errors and recommendations, please refer to: \"%s\".",
			referenceURL,
		)
	}

	result := TestResult{
		Message:         message,
		Success:         success,
		Errors:          errorFeedback,
		Recommendations: recommendationFeedback,
		MoreInfo:        moreInfo,
	}

	return result
}
