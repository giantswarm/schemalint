package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type ExampleExists struct{}

func (r ExampleExists) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema)
	for path, annotations := range propertyAnnotationsMap {
		if len(annotations.GetExamples()) == 0 {
			ruleResults.Add(fmt.Sprintf("Property '%s' should provide one or more examples.", path))
		}
	}
	return *ruleResults
}

func (r ExampleExists) GetSeverity() lint.Severity {
	return lint.SeverityRecommendation
}
