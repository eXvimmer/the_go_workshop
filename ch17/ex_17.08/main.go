package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello mux")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", exampleHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
