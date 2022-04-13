package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	names := `{
    "names": ["Mustafa", "Malena", "Maya", "Mustafa", "Mustafa", "Maya"]
  }`

	w.Write([]byte(names))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
