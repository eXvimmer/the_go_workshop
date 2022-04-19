package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Visitor struct {
	Name string
}

type TemplateHandler struct {
	tmp *template.Template
}

func (th TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := Visitor{}
	q := r.URL.Query()
	names, ok := q["name"]
	if ok {
		v.Name = strings.Join(names, ", ")
	}
	th.tmp.Execute(w, v)
}

func NewTemplateHandler(tmpPath string) (*TemplateHandler, error) {
	temp, err := template.ParseFiles(tmpPath)
	if err != nil {
		return nil, err
	}

	return &TemplateHandler{temp}, nil
}

func main() {
	handler, err := NewTemplateHandler("./index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
