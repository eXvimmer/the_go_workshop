package main

import (
	"log"
	"net/http"
	"time"
)

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "superSecretToken" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authorization token not recognized"))
		return
	}

	time.Sleep(3 * time.Second)
	w.Write([]byte("Hello client"))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
