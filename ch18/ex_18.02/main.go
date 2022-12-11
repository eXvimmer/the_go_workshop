package main

import (
	"fmt"
	"net/http"
	// "text/template" // NOTE : to simulate XSS
	"html/template"
)

var content = `
  <html>
    <head>
      <title>My blog</title>
    </head>
    <body>
      <h1>My blog site</h1>
      <h2>Comments</h2>
      {{.Comment}}
      <form action="/" method="POST">
        Add comment: <input type="text" name="comment" />
        <input type="submit" value="Submit" />
      </form>
    </body>
  </html>
`

type input struct {
	Comment string
}

func handler(w http.ResponseWriter, r *http.Request) {
	userInput := &input{
		Comment: r.FormValue("comment"),
	}
	t := template.Must(template.New("test").Parse(content))
	err := t.Execute(w, userInput)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
