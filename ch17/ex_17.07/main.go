package main

import "fmt"

// Add returns the total of two integers added together
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers multiplied together
func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(Add(4, 9))
	fmt.Println(Multiply(11, 9))

	// NOTE: to see documentations, use 'go doc -all'
}
