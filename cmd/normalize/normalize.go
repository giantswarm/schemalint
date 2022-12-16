package normalize

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/cli"
	"github.com/giantswarm/schemalint/pkg/normalize"
)

var (
	NormalizeCmd = &cobra.Command{
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
		Run:          run,
		SilenceUsage: true,
	}

	outputPath     string
	forceOverwrite bool
)

func init() {
	NormalizeCmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "Output file path. If not set, STDOUT will be used.")
	NormalizeCmd.Flags().BoolVar(&forceOverwrite, "force", false, "Force overwriting any existing file when using --output-path/-o.")
}

func run(cmd *cobra.Command, args []string) {

	path := args[0]
	input := readInputOrExit(path)

	output, err := normalize.Normalize(input)
	if err != nil {
		log.Fatalf("Error processing file %s.\nProbably this is not valid JSON.\nDetails: %s", path, err)
	}

	handleOutput(output, outputPath)
}

func readInputOrExit(path string) []byte {
	input, err := os.ReadFile(path)
	if err != nil {
		cli.FatalfErrorMessage("Error reading file %s: %s", path, err)
	}
	return input
}
