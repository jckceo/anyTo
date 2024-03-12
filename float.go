package anyTo

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Float(n interface{}) float64 {
	switch v := n.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case string:
		s := removeUnicodeChars(v)    // Remove spaces, tabs, newlines, etc.
		if strings.Contains(s, ",") { // Remove commas
			s = strings.Replace(s, ".", "", -1) // Remove the period used as thousands separator
			s = strings.Replace(s, ",", ".", 1) // Replace the comma with a period
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("String: %s - Error: %s\n", s, err.Error())
			f = 0
		}
		return f
	}
	return 0.00
}

func TwoDecimal(num float64) float64 {
	return math.Round(num*100) / 100
}
