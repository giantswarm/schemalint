package verify

import (
	"fmt"
	"sort"

	"github.com/fatih/color"

	"github.com/giantswarm/schemalint/v2/pkg/cli"
)

func printOutput(results []TestResult) {
	recommendations := []string{}
	errors := []string{}

	for _, r := range results {
		recommendations = append(recommendations, r.Recommendations...)
		errors = append(errors, r.Errors...)
	}

	sort.Strings(recommendations)
	sort.Strings(errors)

	printRecommendations(recommendations)
	printErrors(errors)
	printSummary(results)
}

func printRecommendations(recommendations []string) {
	printGenericList("Recommendations", recommendations, cli.WarningColor)
}

func printErrors(errors []string) {
	printGenericList("Errors", errors, cli.ErrorColor)
}

func printGenericList(title string, items []string, color *color.Color) {
	if len(items) == 0 {
		return
	}
	fmt.Println()
	color.Printf("%s (%d)\n", title, len(items))
	fmt.Println()
	for _, i := range items {
		fmt.Printf("- %s\n", i)
	}
}

func printSummary(results []TestResult) {
	var summary string
	for _, r := range results {
		if r.Success {
			summary += cli.SprintSuccessMessage(r.Message) + "\n"
		} else {
			summary += cli.SprintErrorMessage(r.Message) + "\n"
		}
	}

	fmt.Println()
	fmt.Println("Verification result")
	fmt.Println()
	fmt.Println(summary)
}
