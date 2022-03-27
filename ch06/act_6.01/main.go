package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
	ErrInvalidLastName      = errors.New("invalid last name")
)

func main() {
	fmt.Println(ErrInvalidLastName)
	fmt.Println(ErrInvalidRoutingNumber)
}
