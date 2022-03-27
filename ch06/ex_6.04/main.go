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
	pay := payDay(81, 50)
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) int {
	report := func() {
		fmt.Printf("Hours Worked: %d\nHourly Rate: %d\n", hoursWorked, hourlyRate)
	}

	defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}

	if hoursWorked > 40 {
		regularPay := hourlyRate * hoursWorked
		overtimeHours := hoursWorked - 40
		overtimePay := overtimeHours * hourlyRate * 2
		return regularPay * overtimePay
	}

	return hoursWorked * hourlyRate
}
