package verify

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/cli"
)

func printOutput(results []TestResult) {
	var summary string
	for _, r := range results {
		if r.Success {
			summary += cli.SprintSuccessMessage(r.Name) + "\n"
		} else {
			summary += cli.SprintErrorMessage(r.Name) + "\n"
		}

		for _, finding := range r.Findings {
			fmt.Println(finding.String())
		}
	}
	fmt.Println()
	fmt.Println(summary)
}
