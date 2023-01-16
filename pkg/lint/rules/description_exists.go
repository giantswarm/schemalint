package rules

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionExists struct{}

func (r DescriptionExists) Verify(schema *jsonschema.Schema) []string {
	return lint.RecurseProperties(schema, checkDescription)
}

func checkDescription(schema *jsonschema.Schema) []string {
	if schema.Description == "" {
		return []string{fmt.Sprintf("Property '%s' should have a description.", schemautils.GetConciseLocation(schema))}
	}
	return []string{}
}

func (r DescriptionExists) GetSeverity() lint.Severity {
	return lint.SeverityRecomendation
}
