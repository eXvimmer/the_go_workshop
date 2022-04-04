package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	in := `firstName, lastName, age
  Mustafa, Hayati, 29
  Malena, Tudi, 26
  Maya, Higa, 23`

	r := csv.NewReader(strings.NewReader(in))

	header := true // to skip headers

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if !header {
			for i, v := range record {
				switch i {
				case 0:
					fmt.Println("first name:", v)
				case 1:
					fmt.Println("last name:", v)
				case 2:
					fmt.Println("age:", v)
				}
			}
		}
		header = false
	}
}
