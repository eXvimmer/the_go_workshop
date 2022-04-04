package main

import (
	"os"
)

func main() {
	f, err := os.OpenFile("junk101.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.Write([]byte("adding stuff.\n")); err != nil {
		panic(err)
	}
}
