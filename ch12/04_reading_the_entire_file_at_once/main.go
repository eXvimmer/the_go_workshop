package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("test.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("file contents:")
	fmt.Println(string(content))

	fmt.Println("\n-----------------------------------------------------")

	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file contents:")
	fmt.Println(string(contents))

	fmt.Println("\n-----------------------------------------------------")
	r := strings.NewReader("no file here")
	c, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("contents of strings.NewReader:")
	fmt.Println(string(c))
}
