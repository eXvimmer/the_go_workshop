package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Saturday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on a weekend")
	default:
		fmt.Println("error, day born not valid")
	}
}
