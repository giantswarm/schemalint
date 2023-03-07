package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type ShouldDisableAdditionalProperties struct{}

func (r ShouldDisableAdditionalProperties) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		if !isAdditionalPropertiesDisabled(schema) {
			ruleResults.Add(fmt.Sprintf("Object '%s' should disable additional properties.", schema.GetHumanReadableLocation()), schema.GetResolvedLocation())
		}
	}

	utils.RecurseObjects(schema, callback)

	return *ruleResults
}

func isAdditionalPropertiesDisabled(schema *schemautils.ExtendedSchema) bool {
	return schema.AdditionalProperties == false
}

func (r ShouldDisableAdditionalProperties) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
