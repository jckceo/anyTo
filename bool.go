package anyTo

import "strings"

func Bool(n interface{}) bool {
	switch v := n.(type) {
	case bool:
		return v
	case string:
		if strings.Contains(v, "true") || strings.Contains(v, "1") {
			return true
		} else if strings.Contains(v, "false") || strings.Contains(v, "0") {
			return false
		}
	case int:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	}
	return false
}
