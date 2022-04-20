package main

import (
	"log"
)

func push(from, to int, in chan bool, out chan int) {
	for i := from; i <= to; i++ {
		<-in
		out <- i
	}
}

func main() {
	out := make(chan int, 100)
	in := make(chan bool, 100)
	s1 := 0

	go push(1, 25, in, out)
	go push(26, 50, in, out)
	go push(51, 75, in, out)
	go push(76, 100, in, out)

	for i := 0; i < 100; i++ {
		in <- true
		i := <-out
		log.Println(i)
		s1 += i
	}

	log.Println(s1)
}
