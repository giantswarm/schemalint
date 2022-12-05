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
		RunE:         validate,
		SilenceUsage: true,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// Main validation command
func validate(cmd *cobra.Command, args []string) error {
	url, err := lint.ToFileURL(args[0])
	if err != nil {
		return err
	}

	err = lint.CompileAuto(url)
	if err != nil {
		return err
	}

	log.Printf("%s: PASS", args[0])

	return nil
}
