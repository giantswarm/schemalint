package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

const maxDescriptionLength = 200
const minDescriptionLength = 50

type DescriptionShouldHaveCorrectLength struct{}

func (r DescriptionShouldHaveCorrectLength) Verify(schema *schemautils.ExtendedSchema) []string {
	return utils.RecursePropertiesWithDescription(schema, checkDescriptionShouldHaveCorrectLength)
}

func checkDescriptionShouldHaveCorrectLength(schema *schemautils.ExtendedSchema) []string {
	if len(schema.Description) > maxDescriptionLength {
		return []string{fmt.Sprintf("Property '%s' description should be less than %d characters.", schema.GetConciseLocation(), maxDescriptionLength)}
	}

	if len(schema.Description) < minDescriptionLength {
		return []string{fmt.Sprintf("Property '%s' description should be more than %d characters.", schema.GetConciseLocation(), minDescriptionLength)}
	}

	return []string{}
}

func (r DescriptionShouldHaveCorrectLength) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
