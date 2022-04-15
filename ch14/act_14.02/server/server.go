package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var names []string // this is our database

type Name struct {
	Name string `json:"name"`
}

type Names struct {
	Names []string `json:"names"`
}

type Resp struct {
	OK bool `json:"ok"`
}

func AddName(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(400)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var data Name
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	names = append(names, data.Name)

	res := Resp{OK: true}
	bts, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Write(bts)
}

func GetNames(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(400)
		return
	}

	res := Names{Names: names}
	data, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

func main() {
	http.HandleFunc("/addName", AddName)
	http.HandleFunc("/", GetNames)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
