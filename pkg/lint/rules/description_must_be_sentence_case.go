package rules

import (
	"fmt"
	"unicode"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionMustBeSentenceCase struct{}

func (r DescriptionMustBeSentenceCase) Verify(schema *jsonschema.Schema) []string {
	return lint.RecursePropertiesWithDescription(schema, checkDescriptionMustBeSentenceCase)
}

func checkDescriptionMustBeSentenceCase(schema *jsonschema.Schema) []string {
	if !unicode.IsUpper(rune(schema.Description[0])) {
		return []string{fmt.Sprintf("Property '%s' description must start with a capital letter.", schemautils.GetConciseLocation(schema))}
	}
	return []string{}
}

func (r DescriptionMustBeSentenceCase) GetSeverity() lint.Severity {
	return lint.SeverityError
}
