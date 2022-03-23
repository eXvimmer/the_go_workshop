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
	// Sunday Weekday = iota + 1
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

func main() {
	dev := Developer{
		individual: Employee{Id: 13, FirstName: "Mustafa", LastName: "Hayati"},
		HourlyRate: 10_000,
	}

	dev.LogHours(Friday, 8)
	dev.LogHours(Monday, 7)

	fmt.Println("Hours worked on Friday: ", dev.WorkWeek[Friday])
	fmt.Println("Hours worked on Monday: ", dev.WorkWeek[Monday])

	th := dev.HoursWorked()

	fmt.Println("Total hours worked this week: ", th)
	fmt.Println("Total money earned this week: ", th*dev.HourlyRate)
}
