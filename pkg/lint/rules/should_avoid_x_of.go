package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type ShouldAvoidXOf struct{}

func (r ShouldAvoidXOf) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	callback := func(schema *schemautils.ExtendedSchema) {
		if schema.AnyOf != nil {
			ruleResults.Add(fmt.Sprintf("Schema at path '%s' should not use anyOf", schema.GetHumanReadableLocation()))
		}
		if schema.OneOf != nil {
			ruleResults.Add(fmt.Sprintf("Schema at path '%s' should not use oneOf", schema.GetHumanReadableLocation()))
		}
	}
	utils.RecurseAll(schema, callback)
	return *ruleResults
}

func (r ShouldAvoidXOf) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
