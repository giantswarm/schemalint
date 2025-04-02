package verify

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/giantswarm/schemalint/v2/pkg/cli"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
	"github.com/giantswarm/schemalint/v2/pkg/verify"
)

func printOutput(results []verify.TestResult) {
	recommendations, errors, moreInfo := aggregateTestResults(results)

	fmt.Println()
	printRecommendations(recommendations)
	printSeperatorIfNonEmpty(recommendations)
	printErrors(errors)
	printSeperatorIfNonEmpty(errors)
	printMoreInfo(moreInfo)
	printSeperatorIfNonEmpty(moreInfo)
	printSummary(results)
}

func aggregateTestResults(
	results []verify.TestResult,
) ([]verify.Feedback, []verify.Feedback, []string) {

	recommendations := []verify.Feedback{}
	errors := []verify.Feedback{}
	moreInfo := []string{}

	for _, r := range results {
		recommendations = append(recommendations, r.Recommendations...)
		errors = append(errors, r.Errors...)
		if r.MoreInfo != "" {
			moreInfo = append(moreInfo, r.MoreInfo)
		}
	}
	return recommendations, errors, moreInfo
}

func printRecommendations(recommendations []verify.Feedback) {
	printFeedback("Recommendations", recommendations, cli.WarningColor)
}

func printErrors(errors []verify.Feedback) {
	printFeedback("Errors", errors, cli.ErrorColor)
}

func printFeedback(title string, feedback []verify.Feedback, color *color.Color) {
	feedbackGroups := verify.GroupFeedbackByMessage(feedback)
	if len(feedbackGroups) == 0 {
		return
	}

	totalFeedback := 0
	for _, g := range feedbackGroups {
		totalFeedback += len(g)
	}

	_, _ = color.Printf("%s (%d)\n", title, totalFeedback)
	_, _ = fmt.Println()
	for _, feedbackGroup := range feedbackGroups {
		printFeedbackGroup(feedbackGroup)
	}
}

func printFeedbackGroup(feedback []verify.Feedback) {
	if len(feedback) == 0 {
		panic("feedback group must not be empty") // should never happen
	}
	sample := feedback[0]
	fmt.Println(sample.Message)
	if sample.Location != "" {
		for _, f := range feedback {
			fmt.Printf("  - %s\n", schema.ConvertToConciseLocation(f.Location))
		}
	}
	fmt.Println()
}

func printSeperatorIfNonEmpty[T any](s []T) {
	if len(s) > 0 {
		printSeparator()
	}
}

func printSeparator() {
	fmt.Println()
	fmt.Println("---")
	fmt.Println()
}

func printMoreInfo(moreInfo []string) {
	for _, m := range moreInfo {
		fmt.Println(m)
	}
}

func printSummary(results []verify.TestResult) {
	var summary string
	for _, r := range results {
		if r.Success {
			summary += cli.SprintSuccessMessage(r.Message) + "\n"
		} else {
			summary += cli.SprintErrorMessage(r.Message) + "\n"
		}
	}

	fmt.Println("Verification result")
	fmt.Println()
	fmt.Println(summary)
}
