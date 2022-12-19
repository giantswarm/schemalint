package verify

import (
	"os"

	"github.com/giantswarm/schemalint/pkg/cli"
	"github.com/spf13/cobra"
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
