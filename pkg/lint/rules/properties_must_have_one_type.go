package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint"
	"github.com/giantswarm/schemalint/v2/pkg/lint/utils"
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

type PropertiesMustHaveOneType struct{}

func (r PropertiesMustHaveOneType) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		if len(schema.Types) > 1 {
			ruleResults.Add(fmt.Sprintf("Property '%s' has %d types, but must have exactly one.", schema.GetHumanReadableLocation(), len(schema.Types)), schema.GetResolvedLocation())
		}
		if len(schema.Types) == 0 && schema.Ref == nil {
			ruleResults.Add(fmt.Sprintf("Property '%s' must have exactly one type unless '$ref' is used.", schema.GetHumanReadableLocation()), schema.GetResolvedLocation())
		}
	}

	utils.RecurseProperties(schema, callback)
	return *ruleResults
}

func (r PropertiesMustHaveOneType) GetSeverity() lint.Severity {
	return lint.SeverityError
}
