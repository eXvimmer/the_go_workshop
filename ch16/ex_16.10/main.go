package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	in, out chan int
	sbw     int // subworkers
	mtx     *sync.Mutex
}

func (w *Worker) readThem() {
	w.sbw++

	go func() {
		partial := 0

		for i := range w.in {
			partial += i
		}

		w.out <- partial

		w.mtx.Lock()
		w.sbw-- // reduce the counter on the subworker safely
		if w.sbw == 0 {
			close(w.out)
		}
		w.mtx.Unlock()
	}()
}

func (w *Worker) gatherResult() int {
	total := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for o := range w.out {
			total += o
		}

		wg.Done()
	}()

	wg.Wait()
	return total
}

func main() {
	in := make(chan int, 100)
	out := make(chan int)
	mtx := &sync.Mutex{}
	workerCount := 10
	worker := Worker{in: in, out: out, mtx: mtx}

	for i := 1; i <= workerCount; i++ {
		worker.readThem()
	}

	for i := 1; i <= 100; i++ {
		in <- i
	}
	close(in)

	res := worker.gatherResult()
	fmt.Println(res)
}
