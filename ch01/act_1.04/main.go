package main

import (
	"fmt"
)

func main() {
	count := 5
	if count > 0 {
		count = 10
		count++
	}

	fmt.Println(count == 11)
}
