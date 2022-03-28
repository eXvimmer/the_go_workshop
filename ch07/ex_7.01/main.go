// Exercise 7.01: Implementing an Interface

package main

import "fmt"

type Speaker interface {
	Speak() string
}

type person struct {
	name      string
	age       int
	isMarried bool
}

func main() {
	p := person{
		name:      "Mustafa",
		age:       29,
		isMarried: false,
	}

	fmt.Println(p.Speak())
	fmt.Println(p)
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years old).\nMarriage status: %t\n",
		p.name, p.age, p.isMarried)
}

func (p person) Speak() string {
	return "Hi, my name is " + p.name
}
