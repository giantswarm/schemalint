package main

import (
	"log"

	"github.com/giantswarm/schemalint/cmd"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
