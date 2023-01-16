package rules

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

const maxDescriptionLength = 200
const minDescriptionLength = 50

type DescriptionShouldHaveCorrectLength struct{}

func (r DescriptionShouldHaveCorrectLength) Verify(schema *jsonschema.Schema) []string {
	return lint.RecursePropertiesWithDescription(schema, checkDescriptionShouldHaveCorrectLength)
}

func checkDescriptionShouldHaveCorrectLength(schema *jsonschema.Schema) []string {
	if len(schema.Description) > maxDescriptionLength {
		return []string{fmt.Sprintf("Property '%s' description should be less than %d characters.", schemautils.GetConciseLocation(schema), maxDescriptionLength)}
	}

	if len(schema.Description) < minDescriptionLength {
		return []string{fmt.Sprintf("Property '%s' description should be more than %d characters.", schemautils.GetConciseLocation(schema), minDescriptionLength)}
	}

	return []string{}
}

func (r DescriptionShouldHaveCorrectLength) GetSeverity() lint.Severity {
	return lint.SeverityRecomendation
}
