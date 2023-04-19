package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/giantswarm/schemalint/pkg/schema"
)

type MustUseCorrectDialect struct{}

const correctDraft = "https://json-schema.org/draft/2020-12/schema"
const draftKey = "$schema"

func (r MustUseCorrectDialect) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	schemaJson, err := readJson(s.RootFilePath)

	if err != nil {
		// s is already compiled, so it is valid json
		panic(err)
	}

	draft, ok := schemaJson[draftKey].(string)

	if !ok {
		ruleResults.Add(
			fmt.Sprintf(
				"Schema does not specify a draft/dialect, but must use '%s'.",
				correctDraft,
			),
			s.GetResolvedLocation(),
		)
		return *ruleResults
	}

	if draft != correctDraft {
		ruleResults.Add(
			fmt.Sprintf("Schema must use draft/dialect '%s', but uses '%s'.", correctDraft, draft),
			s.GetResolvedLocation(),
		)
	}

	return *ruleResults
}

func (r MustUseCorrectDialect) GetSeverity() Severity {
	return SeverityError
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
