package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	return map[string]string{
		"13": "Mustafa Hayati",
		"89": "Taylor Swift",
		"90": "Katherine Ryan",
	}
}

func getUser(id string) (string, bool) {
	name, ok := getUsers()[id]
	return name, ok
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userId := os.Args[1]
	name, exists := getUser(userId)

	if !exists {
		fmt.Printf("Passed user id (%v) not found.\n", userId)
		for k, v := range getUsers() {
			fmt.Println(" ID:", k, "  Name:", v)
		}
		os.Exit(1)
	}

	fmt.Println("Name:", name)
}
