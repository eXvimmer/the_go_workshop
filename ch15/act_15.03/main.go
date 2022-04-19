package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Person struct {
	Name string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./index.html")
		if err != nil {
			log.Fatal(err)
		}

		p := Person{}

		q := r.URL.Query()
		names, ok := q["name"]
		if ok {
			p.Name = strings.Join(names, ", ")
			temp.Execute(w, p)
			return
		}
		temp.Execute(w, nil)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
