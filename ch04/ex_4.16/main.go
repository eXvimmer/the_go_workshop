package main

import (
	"fmt"
)

type id string

func getIDs() (id, id, id) {
	var (
		id1 id
		id2 id = "id1234"
		id3 id
	)

	id3 = "id1234"

	return id1, id2, id3
}

func main() {
	id1, id2, id3 := getIDs()
	fmt.Println("id1 == id2:  ", id1 == id2)
	fmt.Println("id2 == id3:  ", id2 == id3)
}
