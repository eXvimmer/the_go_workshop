package main

import (
	"fmt"
)

func main() {
	username := "Sir_King_Über"

	fmt.Println("Bytes: ", len(username))
	fmt.Println("Runes: ", len([]rune(username)))

	// NOTE:
	// When using len directly on a string, you get the wrong answer.

	fmt.Println(string(username[:10]))
	fmt.Println()
	fmt.Println(string([]rune(username)[:10]))
}
