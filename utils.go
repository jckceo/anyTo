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

func parseString(s string) string {
	s = removeUnicodeChars(s) // Remove spaces, tabs, newlines, etc.
	if strings.Contains(s, ",") { // Remove commas
		s = strings.Replace(s, ".", "", -1) // Remove the period used as thousands separator
		s = strings.Replace(s, ",", ".", 1) // Replace the comma with a period
	}
	return s
}