package main

import (
	"fmt"
	"os"
)

func getPassedArgs() []string {
	return append([]string{}, os.Args...)[1:]
}

func getLocals(extraLocals []string) []string {
	locals := append([]string{}, "en_US", "fr_FR")
	return append(locals, extraLocals...)
}

func main() {
	fmt.Println("Locals to use: ", getLocals(getPassedArgs()))
}
