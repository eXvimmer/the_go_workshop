package main

import "fmt"

type person struct {
	lname  string
	age    int
	salary float64
}

func main() {
	fname := "Mustafa"
	grades := []int{100, 99, 98}
	states := map[string]string{"NY": "New York", "CA": "California", "TX": "Texas"}
	p := person{
		lname:  "Hayati",
		age:    29,
		salary: 120_000.99,
	}

	fmt.Printf("fname value %#v\n", fname)
	fmt.Printf("grades value %#v\n", grades)
	fmt.Printf("states value %#v\n", states)
	fmt.Printf("p value %#v\n", p)
}
