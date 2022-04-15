package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getDataWithCustomOptionsAndReturnResponse() string {
	client := http.Client{Timeout: 11 * time.Second}
	req, err := http.NewRequest("POST", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "superSecretToken")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	data := getDataWithCustomOptionsAndReturnResponse()
	fmt.Println(data)
}
