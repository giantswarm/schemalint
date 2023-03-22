package normalize

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/v2/pkg/cli"
	"github.com/giantswarm/schemalint/v2/pkg/normalize"
)

type runner struct {
	flag *flag
}

func (r *runner) run(cmd *cobra.Command, args []string) {
	path := args[0]

	input := readInputOrExit(path)

	output, err := normalize.Normalize(input)
	if err != nil {
		log.Fatalf("Error processing file %s.\nProbably this is not valid JSON.\nDetails: %s", path, err)
	}

	handleOutput(output, r.flag.outputPath)
}

func readInputOrExit(path string) []byte {
	input, err := os.ReadFile(path)
	if err != nil {
		cli.FatalfErrorMessage("Error reading file %s: %s", path, err)
	}
	return input
}
