package normalize

import (
	"fmt"
	"os"

	"github.com/giantswarm/schemalint/v2/pkg/cli"
)

func handleOutput(output []byte, outputPath string) {
	if outputPath != "" {
		writeOutputOrExit(output, outputPath)
	} else {
		printOutput(output)
	}

}

func writeOutputOrExit(output []byte, outputPath string) {
	if checkIsUnchanged(output, outputPath) {
		fmt.Printf("Nothing to do, '%s' is already normalized.\n", outputPath)
		return
	}

	err := os.WriteFile(outputPath, output, 0600)
	if err != nil {
		cli.FatalfErrorMessage("Error writing to file: %s", err)
	}
	fmt.Printf("Normalized output written to %s.\n", outputPath)
}

func checkIsUnchanged(output []byte, outputPath string) bool {
	currentFileContent, _ := os.ReadFile(outputPath)
	return string(output) == string(currentFileContent)
}

func printOutput(output []byte) {
	fmt.Print(string(output))
}
