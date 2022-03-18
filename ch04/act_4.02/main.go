package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error, at least 2 arguments are needed")
		os.Exit(1)
	}

	users := map[string]string{
		"305": "Malena",
		"013": "Mustafa",
		"204": "Maya",
		"073": "Mia",
		"647": "Melina",
	}

	fmt.Println("hi,", users[os.Args[1]])
}
