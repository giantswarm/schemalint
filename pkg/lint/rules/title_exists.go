package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type TitleExists struct{}

func (r TitleExists) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema)
	for path, annotations := range propertyAnnotationsMap {
		if annotations.GetTitle() == "" {
			ruleResults.Add(fmt.Sprintf("Property '%s' must have a title.", path))
		}
	}
	return *ruleResults
}

func (r TitleExists) GetSeverity() lint.Severity {
	return lint.SeverityError
}
