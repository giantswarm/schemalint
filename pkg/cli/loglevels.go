package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

const (
	errorPrefix   = "[ERROR] "
	successPrefix = "[SUCCESS] "
)

var (
	SuccessColor *color.Color = color.New(color.FgGreen)
	ErrorColor *color.Color = color.New(color.FgRed).Add(color.Bold)
	WarningColor *color.Color = color.New(color.FgYellow)
)

func PrintErrorMessage(message string) {
	fullMessage := ErrorColor.Sprint(errorPrefix) + message
	fmt.Fprintln(os.Stderr, fullMessage)
}

func SprintErrorMessage(message string) string {
	return ErrorColor.Sprint(errorPrefix) + message
}

func SprintfErrorMessage(format string, a ...interface{}) string {
	return ErrorColor.Sprint(errorPrefix) + fmt.Sprintf(format, a...)
}

func FatalErrorMessage(message string) {
	log.Fatal(SprintErrorMessage(message))
}

func FatalfErrorMessage(format string, a ...interface{}) {
	log.Fatal(SprintfErrorMessage(format, a...))
}

func SprintSuccessMessage(message string) string {
	return SuccessColor.Sprint(successPrefix) + message
}
