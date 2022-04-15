package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var url = "http://localhost:8080"

type Name struct {
	Name string `json:"name"`
}

type Names struct {
	Names []string `json:"names"`
}

type Resp struct {
	OK bool `json:"ok"`
}

func getDataAndParseResponse() []string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	names := Names{}
	err = json.Unmarshal(data, &names)
	if err != nil {
		log.Fatal(err)
	}

	return names.Names
}

func addNameAndParseResponse(theName string) error {
	name := Name{Name: theName}
	nameBytes, err := json.Marshal(name)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(fmt.Sprintf("%s/addName", url), "text/json",
		bytes.NewReader(nameBytes))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp := Resp{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		log.Fatal(err)
	}

	if !resp.OK {
		return errors.New("response not ok")
	}

	return nil
}

func main() {
	err := addNameAndParseResponse("Mustafa")
	if err != nil {
		log.Fatal(err)
	}

	err = addNameAndParseResponse("Malena")
	if err != nil {
		log.Fatal(err)
	}
	names := getDataAndParseResponse()
	for _, name := range names {
		fmt.Println(name)
	}
}
