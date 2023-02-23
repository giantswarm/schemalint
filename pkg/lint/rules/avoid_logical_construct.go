package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type AvoidLogicalConstruct struct{}

func (r AvoidLogicalConstruct) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		if schema.If != nil || schema.Then != nil || schema.Else != nil {
			ruleResults.Add(fmt.Sprintf("Schema must not use logical constructs (if, then, else). Found at '%s'.", schema.GetHumanReadableLocation()))
		}
	}

	utils.RecurseAll(schema, callback)

	return ruleResults
}

func (r AvoidLogicalConstruct) GetSeverity() lint.Severity {
	return lint.SeverityError
}
