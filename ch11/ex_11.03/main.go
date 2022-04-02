package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonData := []byte(`
{
    "id": 8,
    "lname": "Hayati",
    "fname": "Mustafa",
    "IsEnrolled": true,
    "grades":[100,94,98,100],
    "class": 
        {
            "coursename": "Japanese Lit",
            "coursenum": 101,
            "coursehours": 3
        }
}	
`)

	if !json.Valid(jsonData) {
		fmt.Printf("invalid json data %s\n", jsonData)
		os.Exit(1)
	}

	var v interface{}

	err := json.Unmarshal(jsonData, &v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := v.(map[string]interface{})

	for k, v := range data {
		switch value := v.(type) {
		case string:
			fmt.Println("(string):", k, value)
		case bool:
			fmt.Println("(bool):", k, value)
		case float64:
			fmt.Println("(float64):", k, value)
		case []interface{}:
			fmt.Println("(slice):", k, value)
			for i, j := range value {
				fmt.Println("    ", i, j)
			}
		default:
			fmt.Println("(unknown):", k, value)
		}
	}
}
