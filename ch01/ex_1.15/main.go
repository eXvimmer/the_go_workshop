package main

import (
	"fmt"
)

func add5Value(c int) {
	c += 5

	fmt.Println("add5Value: ", c)
}

func add5Point(c *int) {
	*c += 5
	fmt.Println("add5Point: ", *c)
}

func main() {
	var count int

	add5Value(count)
	fmt.Println("add5Value post: ", count)

	add5Point(&count)
	fmt.Println("add5Point post: ", count)
}
