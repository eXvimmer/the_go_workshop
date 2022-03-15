package main

import (
	"fmt"
)

func main() {
	var (
		firstName string = "Mustafa"
		lastName  string = "Hayati"
		age       uint8  = 29
	)
	isAllergicToPeanut := false
	fmt.Println(firstName, lastName, age, isAllergicToPeanut)
}
