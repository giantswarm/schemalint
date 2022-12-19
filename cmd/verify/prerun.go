package verify

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/cli"
)

func (r *runner) preRun(cmd *cobra.Command, args []string) {
	errors := r.flag.validate()
	if len(errors) > 0 {
		for _, e := range errors {
			cli.PrintErrorMessage(e.Error())
		}
		os.Exit(1)
	}
}
