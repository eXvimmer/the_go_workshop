package main

import "fmt"

func main() {
	finished := make(chan struct{})
	names := []string{"Mustafa"}

	go func() {
		// write at this routine
		names = append(names, "Malena")
		names = append(names, "Maya")
		finished <- struct{}{}
	}()

	// and here we're reading in the main routine, so there's a race condition
	// for _, name := range names {
	// 	fmt.Println(name)
	// }
	// <-finished

	// here, we're waiting for the other goroutine to finish its job, so no race
	// condition
	<-finished
	for _, name := range names {
		fmt.Println(name)
	}

	// NOTE: run with `go run --race main.go`
}
