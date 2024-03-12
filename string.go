package anyTo

import "fmt"

func String(n interface{}) string {
	return fmt.Sprintf("%v", n)
}
