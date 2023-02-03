package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type ArraysMustHaveItems struct{}

func (r ArraysMustHaveItems) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		if !hasItems(schema) {
			ruleResults.Add(fmt.Sprintf("Array '%s' must specify the schema of its items.", schema.GetHumanReadableLocation()))
		}
	}

	utils.RecurseArrays(schema, callback)

	return *ruleResults
}

func hasItems(schema *schemautils.ExtendedSchema) bool {
	return schema.Items2020 != nil || schema.Items != nil
}

func (r ArraysMustHaveItems) GetSeverity() lint.Severity {
	return lint.SeverityError
}
