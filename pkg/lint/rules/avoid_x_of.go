package rules

import (
	"fmt"
	"strings"

	"github.com/giantswarm/schemalint/pkg/lint"
	"github.com/giantswarm/schemalint/pkg/lint/utils"
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type AvoidXOf struct{}

func (r AvoidXOf) Verify(schema *schemautils.ExtendedSchema) lint.RuleResults {
	ruleResults := &lint.RuleResults{}
	callback := func(schema *schemautils.ExtendedSchema) {
		if schema.AnyOf != nil {
			permitted, error := isPermittedUsage(schema.GetAnyOf())
			if !permitted {
				ruleResults.Add(fmt.Sprintf(
					"Schema at path '%s' must only use anyOf for one of the following purposes:\n%s",
					schema.GetHumanReadableLocation(),
					error,
				))
			}
		}
		if schema.OneOf != nil {
			permitted, error := isPermittedUsage(schema.GetOneOf())
			if !permitted {
				ruleResults.Add(fmt.Sprintf(
					"Schema at path '%s' must only use oneOf for one of the following purposes:\n%s",
					schema.GetHumanReadableLocation(),
					error,
				))
			}
		}
	}
	utils.RecurseAll(schema, callback)
	return *ruleResults
}

func (r AvoidXOf) GetSeverity() lint.Severity {
	return lint.SeverityError
}

func isPermittedUsage(schemas []*schemautils.ExtendedSchema) (bool, string) {
	permittedAsValidationConstraints, containedIllegalKeywords := isForValidationConstraints(schemas)
	permittedAsDeprecation := isForDeprecation(schemas)

	permitted := permittedAsValidationConstraints || permittedAsDeprecation
	if permitted {
		return true, ""
	}

	validationConstraintsMessage := fmt.Sprintf("Validation constraints: The subschemas cannot be used as validation constraints because they contain the following illegal keywords: %s.", strings.Join(containedIllegalKeywords, ", "))
	permitedAsDeprecationMessage := "Deprecation: The subschemas can only be used for deprecation if exactly one subschema is not deprecated and all others are deprecated."

	return false, fmt.Sprintf("\t- %s\n\t- %s", validationConstraintsMessage, permitedAsDeprecationMessage)
}

// each subschema only defines constraints for the validation of the payload
func isForValidationConstraints(schemas []*schemautils.ExtendedSchema) (bool, []string) {
	containedIllegalKeywords := map[string]bool{}

	for _, schema := range schemas {
		if schema.Types != nil {
			containedIllegalKeywords["type"] = true
		}

		if schema.Title != "" {
			containedIllegalKeywords["title"] = true
		}
		if schema.Description != "" {
			containedIllegalKeywords["description"] = true
		}
		if schema.Examples != nil {
			containedIllegalKeywords["examples"] = true
		}

		if schema.Properties != nil {
			containedIllegalKeywords["properties"] = true
		}
		if schema.AdditionalProperties != nil {
			containedIllegalKeywords["additionalProperties"] = true
		}
		if schema.PatternProperties != nil {
			containedIllegalKeywords["patternProperties"] = true
		}

		if schema.Items != nil || schema.Items2020 != nil {
			containedIllegalKeywords["items"] = true
		}
		if schema.AdditionalItems != nil {
			containedIllegalKeywords["additionalItems"] = true
		}
	}

	keys := make([]string, 0, len(containedIllegalKeywords))
	for k := range containedIllegalKeywords {
		keys = append(keys, k)
	}

	return len(containedIllegalKeywords) == 0, keys
}

// subschemas can be used for deprecation, if exactly one subschema is not
// deprecated and all others are deprecated
func isForDeprecation(schemas []*schemautils.ExtendedSchema) bool {
	nDeprecated := 0
	nNotDeprecated := 0

	for _, schema := range schemas {
		if schema.Deprecated {
			nDeprecated++
		} else {
			nNotDeprecated++
		}
	}
	return nNotDeprecated == 1 && nDeprecated == len(schemas)-1
}