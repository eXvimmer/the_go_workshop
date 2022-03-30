package payroll

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

func PayDetail(p Payer) {
	name, pay := p.Pay()
	fmt.Printf("name: %s\nPay: $%.2f\n", name, pay)
}
