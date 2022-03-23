package main

import (
	"fmt"
)

func main() {
	f := func(i int) int {
		return i * i
	}

	j := 13
	x := f(j)
	fmt.Printf("The square  of %d is %d\n", j, x)
}
