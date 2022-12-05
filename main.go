package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/giantswarm/schemalint/cmd"

	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	cmd.Execute()
}

func readJSON(path string) (interface{}, error) {
	jsonBlob, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var jsonData interface{}
	err = json.Unmarshal(jsonBlob, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, err
}
