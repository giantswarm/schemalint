package cmd

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/cmd/normalize"
	"github.com/giantswarm/schemalint/cmd/verify"
	"github.com/giantswarm/schemalint/pkg/project"
)

var (
	rootCmd = &cobra.Command{
		Use:          "schemalint",
		Short:        "Validate and normalize JSON schema",
		Args:         cobra.MinimumNArgs(1),
		ArgAliases:   []string{"PATH"},
		SilenceUsage: true,
		Version:      project.Version(),
	}
)

func init() {
	rootCmd.AddCommand(normalize.New())
	rootCmd.AddCommand(verify.New())
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
