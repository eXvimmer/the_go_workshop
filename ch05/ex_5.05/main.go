package main

import "fmt"

func main() {
	i := []int{5, 10, 99}

	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))
}

func sum(s ...int) int {
	i := 0

	for _, v := range s {
		i += v
	}

	return i
}
