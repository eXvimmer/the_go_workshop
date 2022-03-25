package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(wd Weekday, h int) {
	d.WorkWeek[wd] = h
}

func (d *Developer) HoursWorked() int {
	t := 0

	for _, h := range d.WorkWeek {
		t += h
	}

	return t
}

func (d *Developer) PayDay() (int, bool) {
	if d.HoursWorked() > 40 {
		regularPay := 4 * d.HourlyRate
		hoursOver := d.HoursWorked() - 40
		overTimePay := hoursOver * 2 * d.HourlyRate
		return regularPay + overTimePay, true
	}

	return d.HoursWorked() * d.HourlyRate, false
}

func (d *Developer) PayDetails() {
	for i, v := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday: ", v, "hours")
		case 1:
			fmt.Println("Monday: ", v, "hours")
		case 2:
			fmt.Println("Tuesday: ", v, "hours")
		case 3:
			fmt.Println("Wednesday: ", v, "hours")
		case 4:
			fmt.Println("Thursday: ", v, "hours")
		case 5:
			fmt.Println("Friday: ", v, "hours")
		case 6:
			fmt.Println("Saturday: ", v, "hours")
		}
	}
	fmt.Println("\nTotal hours worked this week:", d.HoursWorked())
	pay, overtime := d.PayDay()
	fmt.Println("Pay for this week: $", pay)
	fmt.Println("includes overtime:", overtime)
}

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func main() {
	dev := Developer{
		individual: Employee{Id: 13, FirstName: "Mustafa", LastName: "Hayati"},
		HourlyRate: 120,
	}

	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5))
	fmt.Println()

	dev.LogHours(Friday, 8)
	dev.LogHours(Monday, 7)

	dev.LogHours(Monday, 8)
	dev.LogHours(Tuesday, 10)
	dev.LogHours(Wednesday, 10)
	dev.LogHours(Thursday, 10)
	dev.LogHours(Friday, 6)
	dev.LogHours(Saturday, 8)

	dev.PayDetails()
}
