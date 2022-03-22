package main

import (
	"fmt"
)

func main() {
	s := []string{"Good", "Good", "Bad", "Good", "Good"}

	s = append(s[:2], s[3:]...)
	fmt.Println(s)
}
