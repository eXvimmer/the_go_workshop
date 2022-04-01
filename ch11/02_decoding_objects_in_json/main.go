package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	LastName  string  `json:"lname"`
	FirstName string  `json:"fname"`
	Address   address `json:"address"`
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	data := []byte(`
    {
      "lname": "Hayati",
      "fname": "Mustafa",
      "address": {
        "street": "main",
        "city": "New York",
        "state": "New York",
        "zipcode": 13414124
      }
    }
  `)

	if !json.Valid(data) {
		fmt.Printf("Invalid JSON data: %s\n", data)
		os.Exit(1)
	}

	var p person

	err := json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(p.Address.State, p.Address.City)
	fmt.Println(p.Address.Street, p.Address.ZipCode, p.FirstName, p.LastName)
}
