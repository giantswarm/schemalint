package normalize

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/v2/pkg/cli"
)

func (r *runner) preRun(cmd *cobra.Command, args []string) {
	err := r.flag.validate()
	if err != nil {
		cli.FatalErrorMessage(err.Error())
	}
}
