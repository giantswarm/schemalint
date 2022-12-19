package normalize

import "github.com/spf13/cobra"

type flag struct {
	outputPath     string
	forceOverwrite bool
}

func (f *flag) init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.outputPath, "output-path", "o", "", "Output file path. If not set, STDOUT will be used.")
	cmd.Flags().BoolVar(&f.forceOverwrite, "force", false, "Force overwriting any existing file when using --output-path/-o.")
}
