package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 13
	count3 := &countTemp
	t := &time.Time{}

	fmt.Printf("Coun1: %#v\n", count1)
	fmt.Printf("Coun2: %#v\n", count2)
	fmt.Printf("Coun3: %#v\n", count3)
	fmt.Printf("time: %#v\n", t)
}
