package main

import (
	"fmt"
	"time"
)

var (
	Debug       = false
	LogLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
