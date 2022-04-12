package main

import (
	"log"
	"net/http"
)

type server struct {
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := `{
    "message": "hello, world!"
  }`

	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
