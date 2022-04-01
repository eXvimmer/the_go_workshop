package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	StudentId     int      `json:"id"`
	FirstName     string   `json:"fname"`
	Lastname      string   `json:"lname"`
	MiddleInitial string   `json:"mininitial"`
	IsEnrolled    bool     `json:"enrolled"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func main() {
	data := []byte(`
	    {
	    	"id": 123,
	    	"lname": "Hayati",
	    	"minitial": null,
	    	"fname": "Mustafa",
	    	"enrolled": true,
	    	"classes": [{
	    		"coursename": "Intro to Golang",
	    		"coursenum": 101,
	    		"coursehours": 4
	    	},
		{
	    		"coursename": "Japanese Lit",
	    		"coursenum": 101,
	    		"coursehours": 3
	    	},
		{
	    		"coursename": "French Lit",
	    		"coursenum": 101,
	    		"coursehours": 3
	    	}
	
	]
	    }
	`)

	if !json.Valid(data) {
		fmt.Printf("invalid JSON data: %s\n", data)
		os.Exit(1)
	}

	var s student

	err := json.Unmarshal(data, &s)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n\n", s)

	for _, c := range s.Courses {
		fmt.Println(c.Name, c.Hours, "hours")
	}
}
