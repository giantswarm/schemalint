package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
	"github.com/giantswarm/schemalint/v2/pkg/lint/utils"
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

type ObjectsMustHaveProperties struct{}

func (r ObjectsMustHaveProperties) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		nProperties := len(schema.Properties) + len(schema.PatternProperties)
		_, ok := schema.GetAdditionalProperties().(*schemautils.ExtendedSchema)
		if ok {
			nProperties++
		}

		if nProperties == 0 {
			ruleResults.Add(
				fmt.Sprintf(
					"Object at '%s' must have at least one property.",
					schema.GetHumanReadableLocation(),
				),
				schema.GetResolvedLocation(),
			)
		}
	}

	utils.RecurseObjects(schema, callback)
	return *ruleResults
}

func (r ObjectsMustHaveProperties) GetSeverity() lint.Severity {
	return lint.SeverityError
}
