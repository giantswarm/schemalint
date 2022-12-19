package verify

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	flag := &flag{}

	runner := &runner{
		flag: flag,
	}

	cmd := &cobra.Command{
		Use:          "verify PATH",
		Short:        "Verify the given JSON schema input",
		Args:         cobra.ExactArgs(1),
		ArgAliases:   []string{"PATH"},
		PreRun:       runner.preRun,
		Run:          runner.run,
		SilenceUsage: true,
	}
	return cmd
}
