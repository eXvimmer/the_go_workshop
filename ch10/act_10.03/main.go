package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	duration := end.Sub(start)

	fmt.Println("The execution took exactly ", duration, "seconds!")
}
