package anyTo

import (
	"fmt"
	"math"
	"strconv"
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
		s := parseString(v)
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("String: %s - Error: %s\n", s, err.Error())
			return 0
		}
		return f
	}
	return 0.00
}

func TwoDecimal(num float64) float64 {
	return math.Round(num*100) / 100
}
