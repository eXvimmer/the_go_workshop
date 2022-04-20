package main

import "log"

func greeter(ch chan string) {
	ch <- "Hello"
}

func main() {
	ch := make(chan string)
	go greeter(ch)
	log.Println(<-ch)
}
