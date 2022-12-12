package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/normalize"
)

var (
	normalizeCmd = &cobra.Command{
		Use:     "normalize PATH",
		Short:   "Normalize the given JSON schema input",
		Aliases: []string{"normalise", "norm"},
		Long: `Normalize the given JSON schema input.

By default, the normalized JSON will be printed to STDOUT. Use
--output-path / -o to specify a target path. To overwrite an
existing file, add --force.
`,
		Example: `  schemalint normalize schema.json > normalized.json

  schemalint normalize schema.json -o normalized.json

  schemalint normalize in.json -o in.json --force
`,
		Args:         cobra.ExactArgs(1),
		ArgAliases:   []string{"PATH"},
		Run:          normalizeRun,
		SilenceUsage: true,
	}

	outputPath     string
	forceOverwrite bool
)

func init() {
	normalizeCmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "Output file path. If not set, STDOUT will be used.")
	normalizeCmd.Flags().BoolVar(&forceOverwrite, "force", false, "Force overwriting any existing file when using --output-path/-o.")
}

func normalizeRun(cmd *cobra.Command, args []string) {
	if outputPath != "" {
		_, err := os.Stat(outputPath)
		if !forceOverwrite && !errors.Is(err, os.ErrNotExist) {
			log.Fatal("Error: output file already exists. Apply --force to overwrite.")
		}
	}

	path := args[0]
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file %s: %s", path, err)
	}

	output, err := normalize.Normalize(input)
	if err != nil {
		log.Fatalf("Error processing file %s.\nProbably this is not valid JSON.\nDetails: %s", path, err)
	}

	// Write output.
	if outputPath != "" {
		err := os.WriteFile(outputPath, output, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %s", err)
		} else {
			fmt.Printf("Normalized output written to %s.\n", outputPath)
		}
	} else {
		// Print normalized to STDOUT.
		// Caution: no extra white space must be added.
		fmt.Print(string(output))
	}
}
