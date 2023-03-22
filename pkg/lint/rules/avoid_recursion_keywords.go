package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type AvoidRecursionKeywords struct{}

func (r AvoidRecursionKeywords) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	utils.RecurseAll(schema, func(schema *schemautils.ExtendedSchema) {
		if schema.DynamicAnchor != "" || schema.DynamicRef != nil || schema.RecursiveRef != nil {
			ruleResults.Add(
				fmt.Sprintf(
					"Schema at '%s' must not use recursion keywords (dynamicAnchor, dynamicRef, recursiveRef).",
					schema.GetHumanReadableLocation()),
				schema.GetResolvedLocation(),
			)
		}
	})
	return *ruleResults
}

func (r AvoidRecursionKeywords) GetSeverity() lint.Severity {
	return lint.SeverityError
}
