package main

import (
	"fmt"
)

func typeChecker(v interface{}) string {
	switch v.(type) {
	case int, int64, int32:
		return "int"
	case float32, float64:
		return "float"
	case bool:
		return "bool"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func main() {
	s := []interface{}{
		1, 3.14, true, "Mustafa", struct{}{},
	}

	for _, v := range s {
		fmt.Println(v, "is", typeChecker(v))
	}
}
