package normalize

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type flag struct {
	outputPath     string
	forceOverwrite bool
}

func (f *flag) init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.outputPath, "output-path", "o", "", "Output file path. If not set, STDOUT will be used.")
	cmd.Flags().BoolVar(&f.forceOverwrite, "force", false, "Force overwriting any existing file when using --output-path/-o.")
}

func (f *flag) validate() error {
	if f.outputPath != "" && !f.forceOverwrite {
		return errorIfFileExists(f.outputPath)
	}
	return nil
}

func errorIfFileExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("file %s already exists; use --force to overwrite", path)
	}
	if errors.Is(err, os.ErrNotExist) {
		return errors.New("output file already exists; apply --force to overwrite")
	}
	return nil
}
