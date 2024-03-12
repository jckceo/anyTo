package anyTo

import (
	"fmt"
	"strconv"
)

func Int(n interface{}) int {
	max := int(^uint(0) >> 1)

	switch v := n.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		if v <= uint64(max) {
			return int(v)
		}
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		s := parseString(v)
		f, err := strconv.ParseFloat(removeUnicodeChars(s), 64)
		if err != nil {
			fmt.Printf("String: %s - Error: %s\n", s, err.Error())
			return 0
		}
		return int(f)
	}
	return 0
}
