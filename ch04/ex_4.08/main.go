package main

import (
	"fmt"
	"os"
)

func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("at least %v arguments are needed.\n", minArgs)
		os.Exit(1)
	}

	var args []string

	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	// args = append(args, os.Args...)[1:] // I think this is same as for loop

	return args
}

func findLongest(args []string) string {
	longest := ""

	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}

	return longest
}

func main() {
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was: ", longest)
	} else {
		fmt.Println("there was an error")
		os.Exit(1)
	}
}
