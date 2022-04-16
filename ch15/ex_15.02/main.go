package main

import (
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := `<h1>Hello, world!</h1>`
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/chapter1", func(w http.ResponseWriter, r *http.Request) {
		msg := "Chapter 1"
		w.Write([]byte(msg))
	}) // NOTE: this uses the HandleFunc function

	http.Handle("/", hello{}) // NOTE: this uses the Handle function

	log.Fatal(http.ListenAndServe(":8080", nil))
}
