package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/lint"
)

var (
	rootCmd = &cobra.Command{
		Use:          "schemalint PATH",
		Short:        "Validate whether a file is valid JSON schema",
		Args:         cobra.MinimumNArgs(1),
		ArgAliases:   []string{"PATH"},
		Run:          validate,
		SilenceUsage: true,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// Main validation command
func validate(cmd *cobra.Command, args []string) {
	url, err := lint.ToFileURL(args[0])
	if err != nil {
		log.Fatal(err)
	}

	err = lint.CompileAuto(url)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s: PASS", args[0])
}
