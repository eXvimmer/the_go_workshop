package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	Id        string
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func (d Developer) FullName() string {
	return d.Individual.FirstName + " " + d.Individual.LastName
}

func (d Developer) Pay() (string, float64) {
	return d.FullName(), d.HourlyRate * d.HoursWorkedInYear
}

func (d Developer) ReviewRating() float64 {
	var sum float64 = 0

	for _, v := range d.Review {
		switch v.(type) {
		case string:
			r := v.(string)
			switch strings.ToLower(r) {
			case "excellent":
				sum += 5
			case "good":
				sum += 4
			case "fair":
				sum += 3
			case "poor":
				sum += 2
			case "unsatisfactory":
				sum += 1
			default:
				sum += 0
			}

		case int:
			r := v.(int)
			sum += float64(r)
		}
	}

	return sum / float64(len(d.Review))
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (m Manager) FullName() string {
	return m.Individual.FirstName + " " + m.Individual.LastName
}

func (m Manager) Pay() (string, float64) {
	return m.FullName(), m.Salary + (m.CommissionRate * m.Salary)
}

type Payer interface {
	Pay() (string, float64)
}

func payDetail(p Payer) {
	name, pay := p.Pay()
	fmt.Printf("name: %s\nPay: $%.2f\n", name, pay)
}

func main() {
	d := Developer{
		Individual: Employee{
			Id:        "13",
			FirstName: "Mustafa",
			LastName:  "Hayati",
		},
		HourlyRate:        60,
		HoursWorkedInYear: 2300,
		Review: map[string]interface{}{
			"teamwork":        "Good",
			"skill":           5,
			"work quality":    5,
			"problem solving": 4,
			"communication":   "Poor",
			"dependability":   "Unsatisfactory",
		},
	}

	m := Manager{
		Individual: Employee{
			Id:        "9",
			FirstName: "Malena",
			LastName:  "Tudi",
		},
		Salary:         2_000_000,
		CommissionRate: 0.23,
	}

	fmt.Println(d.FullName(), "got a review rating of", d.ReviewRating())
	payDetail(d)
	payDetail(m)
}
