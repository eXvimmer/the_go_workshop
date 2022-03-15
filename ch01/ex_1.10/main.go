package main

import (
	"fmt"
)

func main() {
	count := 5
	fmt.Println(count)

	count += 5
	fmt.Println(count)

	count++
	fmt.Println(count)

	count--
	fmt.Println(count)

	count -= 5
	fmt.Println(count)

	name := "Mustafa"
	name += " Hayati"
	fmt.Println("Hello,", name)
}
