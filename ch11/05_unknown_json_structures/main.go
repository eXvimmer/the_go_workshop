package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := []byte(`{"checkNum":123,"amount":200,"category":["gift","clothing"]}`)
	var v interface{} // v will be map[string]interface{}
	json.Unmarshal(jsonData, &v)
	fmt.Println(v)
}
