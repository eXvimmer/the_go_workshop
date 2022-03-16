package main

import (
	"fmt"
)

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	var word string
	var count int

	for k, v := range words {
		if v > count {
			count = v
			word = k
		}
	}

	fmt.Println("Most popular word:", word)
	fmt.Println("With a count of:", count)
}
