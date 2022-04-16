package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PageWithCounter struct {
	Counter int    `json:"views"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Counter++

	bts, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bts)
}

func main() {
	hello := PageWithCounter{
		Counter: 0,
		Title:   "Hello World",
		Content: "The title of this book is hello world.",
	}

	ch1 := PageWithCounter{
		Counter: 0,
		Title:   "Chapter 1",
		Content: "This is the first chapter of the book.",
	}

	ch2 := PageWithCounter{
		Counter: 0,
		Title:   "Chapter 2",
		Content: "This is the second chapter of the book.",
	}

	http.HandleFunc("/favicon.ico",
		func(w http.ResponseWriter, r *http.Request) {})
	http.Handle("/chapter1", &ch1)
	http.Handle("/chapter2", &ch2)
	http.Handle("/", &hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
