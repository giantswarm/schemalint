package normalize

import (
	"fmt"
	"os"

	"github.com/giantswarm/schemalint/pkg/cli"
)

func handleOutput(output []byte, outputPath string) {
	if outputPath != "" {
		writeOutputOrExit(output, outputPath)
	} else {
		printOutput(output)
	}

}

func writeOutputOrExit(output []byte, outputPath string) {
	err := os.WriteFile(outputPath, output, 0600)
	if err != nil {
		cli.FatalfErrorMessage("Error writing to file: %s", err)
	}
	fmt.Printf("Normalized output written to %s.\n", outputPath)
}

func printOutput(output []byte) {
	fmt.Print(string(output))
}
