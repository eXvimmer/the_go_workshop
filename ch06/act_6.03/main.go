package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
	ErrInvalidLastName      = errors.New("invalid last name")
)

type directDeposit struct {
	firstName     string
	lastName      string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validRoutingNumber() error {
	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}

	return nil
}

func (d *directDeposit) validLastName() error {
	if len(d.lastName) == 0 {
		return ErrInvalidLastName
	}

	return nil
}

func (d *directDeposit) report() {
	fmt.Printf("Last Name: %s\n", d.lastName)
	fmt.Printf("First Name: %s\n", d.firstName)
	fmt.Printf("Bank Name: %s\n", d.bankName)
	fmt.Printf("Routing Number: %d\n", d.routingNumber)
	fmt.Printf("Account Number: %d\n", d.accountNumber)
}

func main() {
	dp := directDeposit{
		firstName:     "Mustafa",
		lastName:      "Hayati",
		bankName:      "Royal Bank of Khiav",
		routingNumber: 9,
		accountNumber: 13,
	}

	err := dp.validLastName()
	if err != nil {
		fmt.Println(err)
	}

	err = dp.validRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}

	dp.report()
}
