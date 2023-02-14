package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type MustUseCorrectDialect struct{}

const correctDraft = "https://json-schema.org/draft/2020-12/schema"
const draftKey = "$schema"

func (r MustUseCorrectDialect) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	schemaJson, err := readJson(schema.RootFilePath)

	if err != nil {
		// schema is already compiled, so it is valid json
		panic(err)
	}

	draft, ok := schemaJson[draftKey].(string)

	if !ok {
		ruleResults.Add(fmt.Sprintf("Schema does not specify a draft/dialect, but must use '%s'.", correctDraft))
		return *ruleResults
	}

	if draft != correctDraft {
		ruleResults.Add(fmt.Sprintf("Schema must use draft/dialect '%s', but uses '%s'.", correctDraft, draft))
	}

	return *ruleResults
}

func (r MustUseCorrectDialect) GetSeverity() lint.Severity {
	return lint.SeverityError
}

func readJson(filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data map[string]interface{}
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
