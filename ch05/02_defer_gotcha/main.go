package main

import "fmt"

func main() {
	age := 29
	name := "Mustafa"

	// NOTE: when a variable is passed to a deferred function, the variable's
	// value at the time is what will be used in the deferred function. Variables
	// used in the deferred function are the values before it was deferred.
	defer personAge(name, age)

	age *= 2
	fmt.Println("age is ", age)
}

func personAge(s string, a int) {
	fmt.Printf("%s is %d years old\n", s, a)
}
