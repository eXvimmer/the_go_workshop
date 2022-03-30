package main

import (
	"fmt"
	"os"

	p "github.com/exvimmer/act_8.01/payroll"
)

var employeeReview = make(map[string]interface{})

func init() {
	fmt.Println("Welcome to employee pay and performance review")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
}

func init() {
	fmt.Println("Initializing variables")
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
}

func main() {
	d := p.Developer{
		Individual: p.Employee{
			Id:        "13",
			FirstName: "Mustafa",
			LastName:  "Hayati",
		},
		HourlyRate:        60,
		HoursWorkedInYear: 2300,
		Review:            employeeReview,
	}

	m := p.Manager{
		Individual: p.Employee{
			Id:        "9",
			FirstName: "Malena",
			LastName:  "Tudi",
		},
		Salary:         2_000_000,
		CommissionRate: 0.23,
	}

	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p.PayDetail(d)
	p.PayDetail(m)
}
