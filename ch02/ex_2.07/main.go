package main

import (
	"fmt"
	"time"
)

func main() {
	switch dayBorn := time.Saturday; {
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Born on a weekend")
	default:
		fmt.Println("Born on some other day")
	}
}
