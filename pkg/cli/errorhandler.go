package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

const (
	errorPrefix   = "ERROR: "
	successPrefix = "SUCCESS: "
)

var (
	successColor *color.Color = color.New(color.FgGreen)
	failureColor *color.Color = color.New(color.FgRed).Add(color.Bold)
)

func PrintErrorMessage(message string) {
	fullMessage := failureColor.Sprint(errorPrefix) + message
	fmt.Fprintln(os.Stderr, fullMessage)
}

func SprintErrorMessage(message string) string {
	return failureColor.Sprint(errorPrefix) + message
}

func SprintfErrorMessage(format string, a ...interface{}) string {
	return failureColor.Sprint(errorPrefix) + fmt.Sprintf(format, a...)
}

func FatalErrorMessage(message string) {
	log.Fatal(SprintErrorMessage(message))
}

func FatalfErrorMessage(format string, a ...interface{}) {
	log.Fatal(SprintfErrorMessage(format, a...))
}

func SprintSuccessMessage(message string) string {
	return successColor.Sprint(successPrefix) + message
}
