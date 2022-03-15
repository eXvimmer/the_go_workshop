package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool      = false
	LogLevl     string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevl, startUpTime)
}
