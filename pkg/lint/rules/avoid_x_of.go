package rules

import (
	"fmt"

	"github.com/giantswarm/schemalint/v2/pkg/lint/recurse"
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type AvoidXOf struct{}

func (r AvoidXOf) Verify(s *schema.ExtendedSchema) RuleResults {
	ruleResults := &RuleResults{}

	validationConstraintsMessage := fmt.Sprintf(
		"Validation constraints: The subschemas cannot be used as validation constraints because they contain one or more of the following illegal keywords: %s.",
		validationConstraintIllegalKeywords,
	)
	permitedAsDeprecationMessage := "Deprecation: The subschemas can only be used for deprecation if exactly one subschema is not deprecated and all others are deprecated."
	allowedUsages := fmt.Sprintf(
		"1. %s\n2. %s",
		validationConstraintsMessage,
		permitedAsDeprecationMessage,
	)

	callback := func(s *schema.ExtendedSchema) {
		if s.AnyOf != nil {
			permitted := isPermittedUsage(s.GetAnyOf())
			if !permitted {
				ruleResults.Add(fmt.Sprintf(
					"Schema must only use anyOf for one of the following purposes:\n%s",
					allowedUsages,
				), s.GetResolvedLocation())
			}
		}
		if s.OneOf != nil {
			permitted := isPermittedUsage(s.GetOneOf())
			if !permitted {
				ruleResults.Add(fmt.Sprintf(
					"Schema at path '%s' must only use oneOf for one of the following purposes:\n%s",
					s.GetConciseLocation(),
					allowedUsages,
				), s.GetResolvedLocation())
			}
		}
	}
	recurse.RecurseAll(s, callback)
	return *ruleResults
}

func (r AvoidXOf) GetSeverity() Severity {
	return SeverityError
}

func isPermittedUsage(schemas []*schema.ExtendedSchema) bool {
	permittedAsValidationConstraints := isForValidationConstraints(
		schemas,
	)
	permittedAsDeprecation := isForDeprecation(schemas)
	return permittedAsValidationConstraints || permittedAsDeprecation
}

var validationConstraintIllegalKeywords = []string{
	"type",
	"title",
	"description",
	"examples",
	"properties",
	"additionalProperties",
	"patternProperties",
	"items",
	"additionalItems",
}

// each subschema only defines constraints for the validation of the payload
func isForValidationConstraints(schemas []*schema.ExtendedSchema) bool {
	for _, schema := range schemas {
		if schema.Types != nil ||
			schema.Title != "" ||
			schema.Description != "" ||
			schema.Examples != nil ||
			schema.Properties != nil ||
			schema.AdditionalProperties != nil ||
			schema.PatternProperties != nil ||
			schema.Items != nil ||
			schema.Items2020 != nil ||
			schema.AdditionalItems != nil {

			return false
		}
	}
	return true
}

// subschemas can be used for deprecation, if exactly one subschema is not
// deprecated and all others are deprecated
func isForDeprecation(schemas []*schema.ExtendedSchema) bool {
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
