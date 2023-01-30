package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type DescriptionExists struct{}

func (r DescriptionExists) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema)

	for path, annotations := range propertyAnnotationsMap {
		if annotations.GetDescription() == "" {
			ruleResults.Add(fmt.Sprintf("Property '%s' should have a description.", path))
		}
	}

	return *ruleResults
}

func (r DescriptionExists) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
