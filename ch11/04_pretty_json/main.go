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
	p := person{
		LastName:  "Riddle",
		FirstName: "Tom",
		Address: address{
			Street:  "main",
			City:    "London",
			State:   "London",
			ZipCode: 1324,
		},
	}

	notPretty, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pretty, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(notPretty))
	fmt.Println()
	fmt.Println(string(pretty))

}
