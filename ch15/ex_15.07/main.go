package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Visitor struct {
	Name    string
	Surname string
	Age     int
}

type Hello struct {
	tpl *template.Template
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := Visitor{}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(400)
			return
		}

		v.Name = r.Form.Get("name")
		v.Surname = r.Form.Get("surname")
		v.Age, _ = strconv.Atoi(r.Form.Get("age"))
	}
	h.tpl.Execute(w, v)
}

func NewHello(tplPath string) (*Hello, error) {
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return nil, err
	}
	return &Hello{tmpl}, nil
}

func main() {
	hello, err := NewHello("./index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", hello)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./form.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
