package main

import "log"

func readThem(in, out chan string) {
	for i := range in {
		log.Println(i)
	}

	out <- "done"
}

func main() {
	// 0 so we don't see anything other than the strings we send
	log.SetFlags(0)
	in, out := make(chan string), make(chan string)
	go readThem(in, out)

	strs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m"}

	for _, str := range strs {
		in <- str
	}
	close(in)

	// the main function terminates only when it has been notified ("done") that
	// all incoming messages have been sent
	<-out
}
