package anyTo

import (
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
			return 0
		}
		return f
	}
	return 0.00
}

func Float2Decimal(n interface{}) float64 {
	return math.Round(Float(n)*100) / 100
}

func Float4Decimal(n interface{}) float64 {
	return math.Round(Float(n)*10000) / 10000
}
