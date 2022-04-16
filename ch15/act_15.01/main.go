package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type PageWithCounter struct {
	counter int
	content string
	heading string
}

func (p *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.counter += 1

	res := fmt.Sprintf(`
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
  </head>
  <body>
    <h1>%s</h1>
    <p>%s</p>
    <p>%s</p>
  </body>
  </html>
  `,
		p.heading,
		p.heading,
		p.content,
		strconv.Itoa(p.counter),
	)

	w.Write([]byte(res))
}

func main() {
	hello := PageWithCounter{
		counter: 0,
		heading: "Hello, World!",
		content: "This is the page for hello world. Have fun!",
	}

	ch1 := PageWithCounter{
		counter: 0,
		heading: "Chapter 1",
		content: "This is the page for chapter 1. I think you'll like it.",
	}

	ch2 := PageWithCounter{
		counter: 0,
		heading: "Chapter 2",
		content: "This is the page for chapter 2. Enjoy!",
	}

	http.HandleFunc("/favicon.ico",
		func(w http.ResponseWriter, r *http.Request) {})
	http.Handle("/chapter1", &ch1)
	http.Handle("/chapter2", &ch2)
	http.Handle("/", &hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
