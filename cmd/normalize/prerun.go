package normalize

import (
	"github.com/giantswarm/schemalint/pkg/cli"
	"github.com/spf13/cobra"
)

func (r *runner) preRun(cmd *cobra.Command, args []string) {
	err := r.flag.validate()
	if err != nil {
		cli.FatalErrorMessage(err.Error())
	}
}
