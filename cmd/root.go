package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "schemalint",
		Short:        "Validate and normalize JSON schema",
		Args:         cobra.MinimumNArgs(1),
		ArgAliases:   []string{"PATH"},
		SilenceUsage: true,
	}
)

func init() {
	rootCmd.AddCommand(normalizeCmd)
	rootCmd.AddCommand(verifyCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
