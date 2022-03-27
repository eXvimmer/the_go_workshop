package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked")
)

func main() {
	pay := payDay(100, 25)
	fmt.Println(pay)
	pay = payDay(100, 200)
	fmt.Println(pay)
	pay = payDay(60, 50)
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHoursWorked {
				fmt.Printf("hours worked: %d\nerr: %v\n\n", hoursWorked, r)
			}

			if r == ErrHourlyRate {
				fmt.Printf("hourly rate: %d\nerr: %v\n\n", hourlyRate, r)
			}
		}

		fmt.Printf(
			"Pay was calculated based on: \nhours worked: %d\nhourly rate: %d\n",
			hoursWorked, hourlyRate,
		)
	}()

	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}

	if hoursWorked > 40 {
		regularPay := hoursWorked * hourlyRate
		overtimeHours := 40 - hoursWorked
		overtimePay := overtimeHours * hourlyRate * 2
		return overtimePay + regularPay
	}

	return hoursWorked * hourlyRate
}
