package cmd

import (
	"log"
	"os"
	"reflect"

	"github.com/spf13/cobra"

	"github.com/giantswarm/schemalint/pkg/lint"
)

var (
	verifyCmd = &cobra.Command{
		Use:          "verify PATH",
		Short:        "Verify the given JSON schema input",
		Args:         cobra.MinimumNArgs(1),
		ArgAliases:   []string{"PATH"},
		Run:          verifyRun,
		SilenceUsage: true,
	}
)

func verifyRun(cmd *cobra.Command, args []string) {
	path := args[0]

	// Verify JSON schema validity (lint)
	{
		url, err := lint.ToFileURL(path)
		if err != nil {
			log.Fatal(err)
		}
		err = lint.CompileAuto(url)
		if err != nil {
			log.Fatal(err)
		}
	}


	log.Printf("%s: PASS", args[0])
}
