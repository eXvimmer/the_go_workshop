package main

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("Thanks for %s", msg)
	ch <- "Hello, Mustafa"
}

func main() {
	ch := make(chan string)
	go greet(ch)
	ch <- "Hello, Malena"
	log.Println(<-ch)
	log.Println(<-ch)
}
