package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func readFileAndPipe(addr string, out chan int, wg *sync.WaitGroup) {
	file, err := os.Open(addr)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				wg.Done()
				return
			}
			panic(err)
		}
		trimmed := strings.ReplaceAll(str, "\n", "")
		i, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(err)
		}
		out <- i
	}
}

func splitter(in, odd, even chan int, wg *sync.WaitGroup) {
	for i := range in {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
	wg.Done()
}

func sum(in, out chan int, wg *sync.WaitGroup) {
	sum := 0
	for i := range in {
		sum += i
	}

	out <- sum
	wg.Done()
}

func merger(even, odd chan int, wg *sync.WaitGroup, resultFile string) {
	rs, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}

	// 2 because we have 2 channels (even and odd)
	for i := 0; i < 2; i++ {
		select {
		case i := <-even:
			rs.Write([]byte(fmt.Sprintf("Even %d\n", i)))
		case i := <-odd:
			rs.Write([]byte(fmt.Sprintf("Odd %d\n", i)))
		}

	}

	rs.Close()
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	wg2 := &sync.WaitGroup{}
	wg2.Add(4)

	odd := make(chan int)
	even := make(chan int)
	out := make(chan int)
	sumOfOdds := make(chan int)
	sumOfEvens := make(chan int)

	go readFileAndPipe("./first.txt", out, wg)
	go readFileAndPipe("./second.txt", out, wg)

	go splitter(out, odd, even, wg2)
	go sum(even, sumOfEvens, wg2)
	go sum(odd, sumOfOdds, wg2)
	go merger(sumOfEvens, sumOfOdds, wg2, "./result.txt")

	wg.Wait()
	close(out)
	wg2.Wait()
}
