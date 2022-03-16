package main

import "fmt"

func main() {
	names := []string{"Mustafa", "Malena", "Maya", "Mia", "Melina"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
