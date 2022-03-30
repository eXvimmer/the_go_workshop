package payroll

import (
	"errors"
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

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d Developer) fullName() string {
	return d.Individual.FirstName + " " + d.Individual.LastName
}

func (d Developer) Pay() (string, float64) {
	return d.fullName(), d.HourlyRate * d.HoursWorkedInYear
}

func (d Developer) ReviewRating() error {
	var sum float64 = 0

	for _, v := range d.Review {
		rating, err := overallReview(v)
		if err != nil {
			return err
		}
		sum += float64(rating)
	}

	averageRating := sum / float64(len(d.Review))

	fmt.Printf("%s got a review rating of %.2f\n", d.fullName(), averageRating)
	return nil
}

func overallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case string:
		switch strings.ToLower(v) {
		case "excellent":
			return 5, nil
		case "good":
			return 4, nil
		case "fair":
			return 3, nil
		case "poor":
			return 2, nil
		case "unsatisfactory":
			return 1, nil
		default:
			return 0, errors.New("invalid rating: " + v)
		}

	case int:
		return v, nil

	default:
		return 0, errors.New("unknown type")
	}
}

func (m Manager) FullName() string {
	return m.Individual.FirstName + " " + m.Individual.LastName
}

func (m Manager) Pay() (string, float64) {
	return m.FullName(), m.Salary + (m.CommissionRate * m.Salary)
}

func PayDetail(p Payer) {
	name, pay := p.Pay()
	fmt.Printf("name: %s\nPay: $%.2f\n", name, pay)
}
