package normalize

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func (r *runner) preRun(cmd *cobra.Command, args []string) {
	outputPath := r.flag.outputPath

	if r.flag.outputPath != "" && !r.flag.forceOverwrite {
		exitIfOutputExists(outputPath)
	}
}

func exitIfOutputExists(path string) {
	exists, err := fileExists(path)
	if err != nil {
		log.Fatalf("Error checking if file exists: %v", err)
		os.Exit(1)
	}
	if exists {
		log.Fatal("Error: output file already exists. Apply --force to overwrite.")
		os.Exit(1)
	}
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
