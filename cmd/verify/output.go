package verify

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/giantswarm/schemalint/pkg/cli"
	"github.com/giantswarm/schemalint/pkg/verify"
)

func printOutput(results []verify.TestResult) {
	recommendations := []verify.Feedback{}
	errors := []verify.Feedback{}
	moreInfo := []string{}

	for _, r := range results {
		recommendations = append(recommendations, r.Recommendations...)
		errors = append(errors, r.Errors...)
		moreInfo = append(moreInfo, r.MoreInfo)
	}

	printRecommendations(recommendations)
	printErrors(errors)
	printMoreInfo(moreInfo)
	printSummary(results)
}

func printFeedbackGroup(feedback []verify.Feedback) {
	if len(feedback) == 0 {
		panic("feedback group must not be empty") // should never happen
	}
	sample := feedback[0]
	fmt.Println(sample.Message)
	if sample.Location != "" {
		for _, f := range feedback {
			fmt.Printf("  - %s\n", f.Location)
		}
	}
	fmt.Println()
}

func printRecommendations(recommendations []verify.Feedback) {
	printFeedback("Recommendations", recommendations, cli.WarningColor)
}

func printErrors(errors []verify.Feedback) {
	printFeedback("Errors", errors, cli.ErrorColor)
}

func printFeedback(title string, feedback []verify.Feedback, color *color.Color) {
	feedbackGroups := verify.GroupByMessage(feedback)
	if len(feedbackGroups) == 0 {
		return
	}

	totalFeedback := 0
	for _, g := range feedbackGroups {
		totalFeedback += len(g)
	}

	printSeparator()
	color.Printf("%s (%d)\n", title, totalFeedback)
	fmt.Println()
	for _, feedbackGroup := range feedbackGroups {
		printFeedbackGroup(feedbackGroup)
	}
}

func printMoreInfo(moreInfo []string) {
	withoutEmpty := []string{}
	for _, i := range moreInfo {
		if i != "" {
			withoutEmpty = append(withoutEmpty, i)
		}
	}
	if len(withoutEmpty) == 0 {
		return
	}
	printSeparator()
	for _, m := range withoutEmpty {
		fmt.Println(m)
	}
}

func printSummary(results []verify.TestResult) {
	var summary string
	printSeparator()
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

func printSeparator() {
	fmt.Println()
}
