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
	// Count occurrences of commas and periods
	commaCount := strings.Count(s, ",")
	periodCount := strings.Count(s, ".")

	if commaCount > 1 && periodCount == 1 {
		// Case: US format with commas as thousand separators and one period as decimal point
		s = strings.Replace(s, ",", "", -1) // Remove all commas
	} else if periodCount > 1 && commaCount == 1 {
		// Case: European format with periods as thousand separators and one comma as decimal point
		s = strings.Replace(s, ".", "", -1) // Remove all periods
		s = strings.Replace(s, ",", ".", 1) // Replace the comma with a period
	} else if commaCount == 1 && periodCount == 1 {
		// Case: Ambiguous format, need to determine based on positions
		if strings.LastIndex(s, ",") > strings.LastIndex(s, ".") {
			// European format
			s = strings.Replace(s, ".", "", -1) // Remove all periods
			s = strings.Replace(s, ",", ".", 1) // Replace the comma with a period
		} else {
			// US format
			s = strings.Replace(s, ",", "", -1) // Remove all commas
		}
	} else if commaCount > 0 && periodCount == 0 {
		// Case: Only commas, assume European format with comma as decimal separator
		s = strings.Replace(s, ",", ".", 1) // Replace the comma with a period
	} /*else if periodCount > 0 && commaCount == 0 {
		// Case: Only periods, assume US format
		// No action needed, already in correct format
	}*/

	return s
}
