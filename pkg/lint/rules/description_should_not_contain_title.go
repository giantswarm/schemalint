package rules

import (
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionShouldNotContainTitle struct{}

func (r DescriptionShouldNotContainTitle) Verify(schema *jsonschema.Schema) []string {
	return lint.RecursePropertiesWithDescription(schema, checkDescriptionShouldNotContainTitle)
}

func checkDescriptionShouldNotContainTitle(schema *jsonschema.Schema) []string {
	if schema.Title == "" {
		return []string{}
	}

	if strings.Contains(strings.ToLower(schema.Description), strings.ToLower(schema.Title)) {
		return []string{fmt.Sprintf("Property '%s' description should not repeat the title.", schemautils.GetConciseLocation(schema))}
	}

	return []string{}
}

func (r DescriptionShouldNotContainTitle) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
