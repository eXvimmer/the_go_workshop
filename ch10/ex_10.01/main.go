package main

import (
	"fmt"
	"time"
)

func whatsTheClock() string {
	return time.Now().Format(time.ANSIC)
}

func main() {
	fmt.Println(whatsTheClock())
}
