package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/normalize"
)

var (
	normalizeCmd = &cobra.Command{
		Use:          "normalize PATH",
		Short:        "Normalize the given JSON schema input",
		Args:         cobra.MinimumNArgs(1),
		ArgAliases:   []string{"PATH"},
		Run:          normalizeRun,
		SilenceUsage: true,
	}
)

func normalizeRun(cmd *cobra.Command, args []string) {
	path := args[0]
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file %s: %s", path, err)
	}

	output, err := normalize.Normalize(input)
	if err != nil {
		log.Fatalf("Error reading file %s: %s", path, err)
	}

	fmt.Println(string(output))
}
