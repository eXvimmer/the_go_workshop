package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func main() {
	_, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}

	_, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}

	_, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}

}

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, ErrHourlyRate
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, ErrHoursWorked
	}

	if hoursWorked > 40 {
		regularPay := hoursWorked * hourlyRate
		hoursOver := hoursWorked - 40
		overtimePay := hoursOver * hourlyRate * 2
		return regularPay + overtimePay, nil
	}

	return hoursWorked * hourlyRate, nil
}
