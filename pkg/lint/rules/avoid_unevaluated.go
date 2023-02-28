package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type AvoidUnevaluated struct{}

func (r AvoidUnevaluated) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		if schema.UnevaluatedItems != nil {
			ruleResults.Add(fmt.Sprintf("Property '%s' must not use unevaluatedItems.", schema.GetHumanReadableLocation()))
		}
		if schema.UnevaluatedProperties != nil {
			ruleResults.Add(fmt.Sprintf("Property '%s' must not use unevaluatedProperties.", schema.GetHumanReadableLocation()))
		}
	}

	utils.RecurseAll(schema, callback)

	return *ruleResults
}

func (r AvoidUnevaluated) GetSeverity() lint.Severity {
	return lint.SeverityError
}
