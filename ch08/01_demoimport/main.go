package main

import (
	"fmt"
	"msg"
)

// NOTE: putting this package in ~/go/src/msg/msg.go didn't help. So I put it
// at /usr/local/go/src/msg/msg.go (GOROOT) and it worked.

// TODO: remove /usr/local/go/src/msg/msg.go after you finished

func main() {
	fmt.Println("Demo import app")
	msg.Greeting("Mustafa")
}
