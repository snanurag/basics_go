package math

import "fmt"

func min(x, y interface{}) interface{} {
	switch v := x.(type) {
	case int:
		if x.(int) < y.(int) {
			return x
		}
		return y
	case float32:
		if x.(float32) < y.(float32) {
			return x
		}
		return y
	case float64:
		if x.(float64) < y.(float64) {
			return x
		}
		return y
	default:
		fmt.Errorf("Type mismatch x %s -> x : %s and y : %s", v, x, y)
		return nil
	}
}
