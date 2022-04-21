package main

import (
	"context"
	"log"
	"time"
)

func countNumbers(c context.Context, ch chan int) {
	v := 0
	for {
		select {
		case <-c.Done():
			ch <- v
			break
		default:
			time.Sleep(time.Millisecond * 100)
			v++
		}
	}
}

func main() {
	ch := make(chan int)
	c := context.TODO()
	cl, stop := context.WithCancel(c)
	go countNumbers(cl, ch)

	go func() {
		time.Sleep(time.Millisecond * 100 * 3)
		stop()
	}()

	v := <-ch
	log.Println(v)
}
