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
// without a '$ref' to close, is not covered.
func isClosedRef(s *schema.ExtendedSchema) bool {
	unevaluated := s.UnevaluatedProperties
	return s.Ref != nil &&
		unevaluated != nil &&
		unevaluated.Always != nil &&
		!*unevaluated.Always
}
