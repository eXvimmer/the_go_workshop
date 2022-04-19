package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Response struct {
	Greeting string `json:"greeting"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data Request
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	rsp := Response{
		Greeting: fmt.Sprintf("Hello %s %s", data.Name, data.Surname),
	}

	bts, err := json.Marshal(rsp)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Write(bts)
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
