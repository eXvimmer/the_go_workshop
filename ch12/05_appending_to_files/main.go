package main

import (
	"os"
)

func main() {
	f, err := os.OpenFile("junk101.txt", os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()
}
