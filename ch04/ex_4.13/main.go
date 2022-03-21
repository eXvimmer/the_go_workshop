package main

import "fmt"

func getUsers() map[string]string {
	users := map[string]string{
		"13": "Mustafa Hayati",
		"99": "Margot Robbie",
		"89": "Taylor Swift",
		"4":  "Sean Lock",
	}

	users["47"] = "Bill Burr"

	return users
}

func main() {
	fmt.Println("Users: ", getUsers())
}
