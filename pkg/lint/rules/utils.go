package rules

import (
	"strings"
	"unicode"

	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

func getIllegalCharacterIn(s string, illegalCharacters []string) (containedIllegalCharacters []string) {
	for _, illegalCharacter := range illegalCharacters {
		if strings.Contains(s, illegalCharacter) {
			containedIllegalCharacters = append(containedIllegalCharacters, illegalCharacter)
		}
	}

	return containedIllegalCharacters
}

func containsLeadingOrTrailingSpace(s string) bool {
	return strings.HasPrefix(s, " ") || strings.HasSuffix(s, " ")
}

func stringStartsCapitalized(str string) bool {
	return unicode.IsUpper(rune(str[0]))

}

// isClosedRef reports whether s closes a '$ref'ed object with
// 'unevaluatedProperties: false'.
//
// That is the only correct spelling in draft 2020-12: 'additionalProperties' only
// considers sibling 'properties' and 'patternProperties' (2020-12 core, section
// 10.3.2), never properties pulled in through '$ref', so 'additionalProperties:
// false' next to a '$ref' rejects every field the referenced schema defines.
//
// The check is deliberately narrow -- an 'unevaluatedProperties' subschema, or one
// without a '$ref' to close, is not covered, and neither is a spelling that leaves
// the object open anyway.
func isClosedRef(s *schema.ExtendedSchema) bool {
	if s.Ref == nil {
		return false
	}

	// A sibling 'additionalProperties' that is not 'false' annotates every property
	// as evaluated, so 'unevaluatedProperties: false' is left with nothing to
	// reject and the object stays open.
	if s.AdditionalProperties != nil && s.AdditionalProperties != false {
		return false
	}

	// 'unevaluatedProperties' has no effect on a non-object. A schema without an
	// explicit type takes it from the '$ref' target, so only a declared type that
	// is not 'object' disqualifies.
	if len(s.Types) > 0 && !s.IsObject() {
		return false
	}

	unevaluated := s.UnevaluatedProperties

	return unevaluated != nil &&
		unevaluated.Always != nil &&
		!*unevaluated.Always
}

// isClosedRefTarget reports whether s is the schema a closed '$ref' points at.
//
// recurse visits the target as well, and resolves it to the *same* location as the
// referring schema, so it has to share the isClosedRef exemption -- otherwise a
// rule still fires at the very location the exemption is meant to clear.
func isClosedRefTarget(s *schema.ExtendedSchema) bool {
	return s.Parent != nil && s.Parent.Ref == s.Schema && isClosedRef(s.Parent)
}
