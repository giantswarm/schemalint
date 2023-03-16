package main

import (
	"log"

	"github.com/giantswarm/schemalint/v2/cmd"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	_ = cmd.Execute()
}
