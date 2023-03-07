package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type AvoidRecursion struct{}

func (r AvoidRecursion) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	utils.RecurseAllPre(schema, func(schema *schemautils.ExtendedSchema) {
		if schema.IsSelfReference() {
			ruleResults.Add(
				fmt.Sprintf(
					"Schema at '%s' must not reference itself.",
					schema.GetHumanReadableLocation(),
				),
				schema.GetResolvedLocation(),
			)
		}
	})
	return *ruleResults
}

func (r AvoidRecursion) GetSeverity() lint.Severity {
	return lint.SeverityError
}
