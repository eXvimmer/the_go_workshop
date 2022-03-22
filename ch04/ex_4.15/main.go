package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"13": "Mustafa Hayati",
	"49": "Sean Lock",
	"99": "Taylor Swift",
	"89": "Killua",
}

func deleteUser(id string) {
	delete(users, id)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("user id not passed")
		os.Exit(1)
	}

	userId := os.Args[1]
	deleteUser(userId)
	fmt.Println("Users:", users)
}
