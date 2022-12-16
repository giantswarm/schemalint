package verify

import (
	"fmt"
	"os"

	"github.com/giantswarm/schemalint/pkg/cli"
)

func printOutput(results []TestResult) {
	var summary string
	for _, r := range results {
		if r.Success {
			summary += cli.SprintSuccessMessage(r.Name) + "\n"
		}

		for _, finding := range r.Findings {
			fmt.Fprintln(os.Stderr, finding)
		}
	}
	summary += "\n"

	fmt.Println(summary)
}
