package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type namesData struct {
	Names []string `json:"names"`
}

func getDataAndParseReponse() (int, int) {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	names := namesData{}

	err = json.Unmarshal(data, &names)
	if err != nil {
		log.Fatal(err)
	}

	mustafaCount := 0
	malenaCount := 0

	for _, n := range names.Names {
		if n == "Mustafa" {
			mustafaCount++
		} else if n == "Malena" {
			malenaCount++
		}
	}

	return mustafaCount, malenaCount
}

func main() {
	mustafaCount, malenaCount := getDataAndParseReponse()
	fmt.Println("Mustafa repeated", mustafaCount, "time(s)")
	fmt.Println("Malena repeated", malenaCount, "time(s)")
}
