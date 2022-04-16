package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	names, ok := q["name"]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Hello stranger"))
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello %s!", strings.Join(names, ", "))))
}

func main() {
	http.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
