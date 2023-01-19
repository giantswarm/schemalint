package rules

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

// Check recursively that all properties have a title
type TitleExists struct{}

func (r TitleExists) Verify(schema *jsonschema.Schema) []string {
	return lint.RecurseProperties(schema, checkTitle)
}

func checkTitle(schema *jsonschema.Schema) []string {
	if schema.Title == "" {
		return []string{fmt.Sprintf("Property '%s' must have a title.", schemautils.GetConciseLocation(schema))}
	}
	return []string{}
}

func (r TitleExists) GetSeverity() lint.Severity {
	return lint.SeverityError
}
