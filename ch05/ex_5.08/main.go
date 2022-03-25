package main

import "fmt"

// type accountant func(int, int) int

func main() {
	devSalary := salary(50, 208, developerSalary)
	bossSalary := salary(150_000, 25_000, managerSalary)

	fmt.Printf("Boss Salary: %d\n", bossSalary)
	fmt.Printf("Dev Salary: %d\n", devSalary)
}

// func salary(x, y int, f accountant) int {
func salary(x, y int, f func(int, int) int) int {
	return f(x, y)
}

func developerSalary(hourlyRate, hoursWorked int) int {
	return hourlyRate * hoursWorked
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}
