package cmd

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/normalize"
)

var (
	verifyCmd = &cobra.Command{
		Use:          "verify PATH",
		Short:        "Verify the given JSON schema input",
		Args:         cobra.MinimumNArgs(1),
		ArgAliases:   []string{"PATH"},
		Run:          verifyRun,
		SilenceUsage: true,
	}

	successColor *color.Color = color.New(color.FgGreen)
	failureColor *color.Color = color.New(color.FgRed).Add(color.Bold)
)

// Structure to collect results from different checks
type TestResult struct {
	Name          string
	Success       bool
	ReturnedError error
}

func verifyRun(cmd *cobra.Command, args []string) {
	path := args[0]

	errorCount := 0
	results := []TestResult{}

	// Verify JSON schema validity
	{
		url, err := lint.ToFileURL(path)
		if err == nil {
			err = lint.CompileAuto(url)
		}

		if err != nil {
			errorCount += 1
		}

		result := TestResult{
			Name:          "Is valid JSON schema",
			Success:       err == nil,
			ReturnedError: err,
		}
		results = append(results, result)
	}

	// Verify normalization
	{
		var err error
		var got, want []byte

		got, err = os.ReadFile(path)
		if err == nil {
			want, err = normalize.Normalize(got)
			if err == nil && !reflect.DeepEqual(got, want) {
				err = errors.New("file is not normalized")
			}
		}

		if err != nil {
			errorCount += 1
		}

		result := TestResult{
			Name:          "Is normalized",
			Success:       err == nil,
			ReturnedError: err,
		}
		results = append(results, result)
	}

	// Preparing our output.
	var output string
	for _, r := range results {
		resultString := successColor.Sprint("SUCCESS")
		if !r.Success {
			resultString = failureColor.Sprint("FAILED")
		}

		// Collect summary for the end.
		output += "\n" + fmt.Sprintf("%s: %s", r.Name, resultString)

		// Print errors to STDERR immediately.
		if r.ReturnedError != nil {
			os.Stderr.Write([]byte("Error: " + r.ReturnedError.Error() + "\n"))
		}
	}
	output += "\n"

	os.Stderr.Write([]byte(output))

	if errorCount > 0 {
		os.Exit(1)
	}
}
