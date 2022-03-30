package main

import (
	"fmt"
	"github.com/exvimmer/act_8.01/payroll"
)

func main() {
	d := payroll.Developer{
		Individual: payroll.Employee{
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

	m := payroll.Manager{
		Individual: payroll.Employee{
			Id:        "9",
			FirstName: "Malena",
			LastName:  "Tudi",
		},
		Salary:         2_000_000,
		CommissionRate: 0.23,
	}

	fmt.Println(d.FullName(), "got a review rating of", d.ReviewRating())
	payroll.PayDetail(d)
	payroll.PayDetail(m)
}
