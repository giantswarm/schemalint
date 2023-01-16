package rules

import (
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionMustNotContainLineBreaks struct{}

func (r DescriptionMustNotContainLineBreaks) Verify(schema *jsonschema.Schema) []string {
	return lint.RecursePropertiesWithDescription(schema, checkDescriptionDoesNotContainLineBreaks)
}

func checkDescriptionDoesNotContainLineBreaks(schema *jsonschema.Schema) []string {
	if containsLineBreaks(schema.Description) {
		return []string{fmt.Sprintf("Property '%s' description must not contain line breaks.", schemautils.GetConciseLocation(schema))}
	}
	return []string{}
}

func containsLineBreaks(s string) bool {
	return strings.Contains(s, "\n")
}

func (r DescriptionMustNotContainLineBreaks) GetSeverity() lint.Severity {
	return lint.SeverityError
}
