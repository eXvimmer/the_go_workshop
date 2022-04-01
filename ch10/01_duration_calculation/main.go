package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("The script started at ", start)
	sum := 0
	for i := 0; i < 10_000_000_000; i++ {
		sum += i
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("The script completed at:", end)
	fmt.Println("The task took", duration.Hours(), "hour(s) to complete")
	fmt.Println("The task took", duration.Minutes(), "minute(s) to complete")
	fmt.Println("The task took", duration.Seconds(), "second(s) to complete")
	fmt.Println("The task took", duration.Nanoseconds(), "nanosecond(s) to complete")
}
