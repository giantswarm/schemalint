package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionExists struct{}

func (r DescriptionExists) Verify(schema *schemautils.ExtendedSchema) []string {
	return lint.RecurseProperties(schema, checkDescriptionExists)
}

func checkDescriptionExists(schema *schemautils.ExtendedSchema) []string {
	if schema.Description == "" {
		return []string{fmt.Sprintf("Property '%s' should have a description.", schema.GetConciseLocation())}
	}
	return []string{}
}

func (r DescriptionExists) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
