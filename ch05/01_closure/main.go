package main

import "fmt"

func main() {
	increment := incrementer()

	fmt.Println(increment())
	fmt.Println(increment())
}

func incrementer() func() int {
	i := 0

	return func() int {
		i += 1
		return i
	}
}
