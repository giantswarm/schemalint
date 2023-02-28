package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type AdheresToCommonSchemaStructureRequirements struct{}

func (r AdheresToCommonSchemaStructureRequirements) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}

	requiredProperties := getRequiredProperties()

	schemaProperties := schema.GetProperties()
	for _, requiredProperty := range requiredProperties {
		if property, ok := schemaProperties[requiredProperty.Name]; ok {
			if !property.IsType(requiredProperty.Type) {
				ruleResults.Add(fmt.Sprintf("Property '%s' must be of type '%s'.", requiredProperty.Name, requiredProperty.Type))
			}
		} else {
			ruleResults.Add(fmt.Sprintf("Property '%s' must be present.", requiredProperty.Name))
		}

	}

	return *ruleResults
}

type propertyRequirement struct {
	Name string
	Type string
}

func getRequiredProperties() []propertyRequirement {
	requiredProperties := []propertyRequirement{
		{
			Name: "metadata",
			Type: "object",
		},
		{
			Name: "connectivity",
			Type: "object",
		},
		{
			Name: "controlPlane",
			Type: "object",
		},
		{
			Name: "nodePools",
			Type: "array",
		},
	}

	return requiredProperties
}

func (r AdheresToCommonSchemaStructureRequirements) GetSeverity() lint.Severity {
	return lint.SeverityError
}
