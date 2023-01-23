package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

// Check recursively that all properties have a title
type TitleExists struct{}

func (r TitleExists) Verify(schema *schemautils.ExtendedSchema) []string {
	return utils.RecurseProperties(schema, checkTitle)
}

func checkTitle(schema *schemautils.ExtendedSchema) []string {
	if schema.Title == "" {
		return []string{fmt.Sprintf("Property '%s' must have a title.", schema.GetConciseLocation())}
	}
	return []string{}
}

func (r TitleExists) GetSeverity() lint.Severity {
	return lint.SeverityError
}
