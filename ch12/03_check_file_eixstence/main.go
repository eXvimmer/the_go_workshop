package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Stat("junk.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("junk.txt: file doesn't exist")
			fmt.Println(file)
		}
	}

	file, err = os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("test.txt: file doesn't exist")
		}
	}
	fmt.Printf("file name: %s\nIsDir: %t\nModTime: %v\nMode: %v\nSize: %d\n",
		file.Name(), file.IsDir(), file.ModTime(), file.Mode(), file.Size())
}
