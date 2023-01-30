package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type TitleMustNotContainParentsTitle struct{}

func (r TitleMustNotContainParentsTitle) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	propertyAnnotationsMap := utils.BuildPropertyAnnotationsMap(schema).WhereTitlesExist()

	for path, annotations := range propertyAnnotationsMap {
		title := annotations.GetTitle()
		parentTitle := propertyAnnotationsMap.GetParentAnnotations(path).GetTitle()

		if parentTitle == "" {
			continue
		}

		if strings.Contains(strings.ToLower(title), strings.ToLower(parentTitle)) {
			ruleResults.Add(fmt.Sprintf("Property '%s' title must not contain parent title '%s'.", path, parentTitle))
		}
	}
	return *ruleResults
}

func (r TitleMustNotContainParentsTitle) GetSeverity() lint.Severity {
	return lint.SeverityError
}
