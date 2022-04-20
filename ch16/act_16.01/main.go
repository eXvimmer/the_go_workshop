// Mutex is short for "mutual exclusion" and is essentially a way to stop all
// the routines, run the code in one, and then carry on with the concurrent
// code.
package main

import (
	"fmt"
	"log"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, res *string, m *sync.Mutex) {
	for i := from; i <= to; i++ {
		m.Lock()
		*res = fmt.Sprintf("%s|%d|", *res, i)
		m.Unlock()
	}

	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	res := ""

	wg.Add(4)
	go sum(1, 25, wg, &res, m)
	go sum(26, 50, wg, &res, m)
	go sum(51, 75, wg, &res, m)
	go sum(76, 100, wg, &res, m)

	wg.Wait()
	log.Println(res)
}
