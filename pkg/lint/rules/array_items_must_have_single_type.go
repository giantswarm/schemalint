package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type ArrayItemsMustHaveSingleType struct{}

func (r ArrayItemsMustHaveSingleType) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	callback := func(schema *schemautils.ExtendedSchema) {
		if containedKeywords := getContainedIllegalArrayKeywords(schema); len(containedKeywords) != 0 {
			ruleResults.Add(
				fmt.Sprintf(
					"Array at '%s' must not use keyword(s): %s.",
					schema.GetHumanReadableLocation(),
					strings.Join(containedKeywords, ", "),
				),
				schema.GetResolvedLocation(),
			)
		}
	}

	utils.RecurseArrays(schema, callback)

	return *ruleResults
}

func getContainedIllegalArrayKeywords(schema *schemautils.ExtendedSchema) []string {
	containedKeywords := []string{}
	if schema.AdditionalItems != nil {
		containedKeywords = append(containedKeywords, "additionalItems")
	}
	if schema.Contains != nil {
		containedKeywords = append(containedKeywords, "contains")
	}
	if schema.PrefixItems != nil {
		containedKeywords = append(containedKeywords, "prefixItems")
	}
	return containedKeywords
}

func (r ArrayItemsMustHaveSingleType) GetSeverity() lint.Severity {
	return lint.SeverityError
}
