package normalize

import (
	"github.com/spf13/cobra"
)

const (
	longDescription = `Normalize the given JSON schema input.

By default, the normalized JSON will be printed to STDOUT. Use
--output-path / -o to specify a target path. To overwrite an
existing file, add --force.
`
	example = `  schemalint normalize schema.json > normalized.json
  schemalint normalize schema.json -o normalized.json
  schemalint normalize in.json -o in.json --force
`
)

func New() *cobra.Command {
	flag := &flag{}

	runner := &runner{
		flag: flag,
	}

	cmd := &cobra.Command{
		Use:          "normalize PATH",
		Short:        "Normalize the given JSON schema input",
		Aliases:      []string{"normalise", "norm"},
		Long:         longDescription,
		Example:      example,
		Args:         cobra.ExactArgs(1),
		ArgAliases:   []string{"PATH"},
		Run:          runner.run,
		PreRun:       runner.preRun,
		SilenceUsage: true,
	}

	flag.init(cmd)
	return cmd
}
