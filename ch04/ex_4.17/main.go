package main

import (
	"fmt"
)

type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

func getUsers() []user {
	u1 := user{
		name:    "Mustafa Hayati",
		age:     29,
		balance: 189.0,
		member:  false,
	}

	u2 := user{
		"Taylor Swift",
		32,
		498_349_259.99,
		true,
	}
	u3 := user{
		name: "Margot Robbie",
		age:  31,
	}

	var u4 user
	u4.name = "Maya Higa"
	u4.age = 23
	u4.balance = 2_000_000.00
	u4.member = false

	return []user{u1, u2, u3, u4}
}

func main() {
	users := getUsers()

	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}
}
