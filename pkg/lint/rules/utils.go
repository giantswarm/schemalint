package rules

import (
	"strings"
	"unicode"
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
