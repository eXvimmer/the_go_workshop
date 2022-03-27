package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := "good-bye"
	message(msg)
	fmt.Println("This line will never get printed")
}

func message(s string) {
	if s == "good-bye" {
		panic(errors.New("Something went wrong"))
	}
}
