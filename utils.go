package anyTo

import (
	"strings"
	"unicode"
)

func removeUnicodeChars(input string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsSpace(r) && !unicode.IsSymbol(r) {
			return r
		}
		return -1 // remove character
	}, input)
}
