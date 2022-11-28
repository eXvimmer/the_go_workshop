package main

import (
	"log"
	"time"
)

func push(from, to int, ch chan int) {
	for i := from; i <= to; i++ {
		ch <- i
		time.Sleep(time.Microsecond) // let another routine pick up the work
	}
}

func main() {
	s1 := 0
	ch := make(chan int, 100)
	go push(1, 25, ch)
	go push(26, 50, ch)
	go push(51, 75, ch)
	go push(76, 100, ch)

	for i := 0; i < 100; i++ {
		c := <-ch
		log.Println(c)
		s1 += c
	}

	log.Println(s1)
}
