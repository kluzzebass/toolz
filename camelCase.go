package toolz

import (
	"strings"
	"unicode"
)

// toCamelCase converts a PascalCase string to camelCase.
func ToCamelCase(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}

	// Lowercase the first letter
	runeArr := []rune(s)
	runeArr[0] = unicode.ToLower(runeArr[0])

	return string(runeArr)
}
