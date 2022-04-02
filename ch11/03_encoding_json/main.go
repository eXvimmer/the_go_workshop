package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type greeting struct {
	SomeMessage string
}

type book struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub"`
	Author        string `json:"author"`
	CoAuthor      string `json:"coauthor"`
}

type book2 struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub,omitempty"`
	Author        string `json:"author"`
	CoAuthor      string `json:"coauthor,omitempty"`
}

type book3 struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:",omitempty"`
	Author        string `json:",omitempty"`
	CoAuthor      string `json:"-"`
}

func main() {
	var v greeting
	v.SomeMessage = "Marshall me!"

	data, err := json.Marshal(v)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", data)

	var b = book{
		ISBN:   "13429something",
		Title:  "A great books",
		Author: "Mustafa Hayati",
	}

	data, err = json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", data)

	var b2 = book2{
		ISBN:   "13429something",
		Title:  "A great books",
		Author: "Mustafa Hayati",
	}

	data, err = json.Marshal(b2)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", data)

	var b3 = book3{
		ISBN:     "13429something",
		Title:    "A great books",
		Author:   "Mustafa Hayati",
		CoAuthor: "Can't see me",
	}

	data, err = json.Marshal(b3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", data)

}
