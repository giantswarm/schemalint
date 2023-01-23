package rules

import (
	"fmt"
	"unicode"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionMustBeSentenceCase struct{}

func (r DescriptionMustBeSentenceCase) Verify(schema *schemautils.ExtendedSchema) []string {
	return utils.RecursePropertiesWithDescription(schema, checkDescriptionMustBeSentenceCase)
}

func checkDescriptionMustBeSentenceCase(schema *schemautils.ExtendedSchema) []string {
	if !unicode.IsUpper(rune(schema.Description[0])) {
		return []string{fmt.Sprintf("Property '%s' description must start with a capital letter.", schema.GetConciseLocation())}
	}
	return []string{}
}

func (r DescriptionMustBeSentenceCase) GetSeverity() lint.Severity {
	return lint.SeverityError
}
