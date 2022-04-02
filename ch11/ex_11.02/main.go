package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"mname,omitempty"`
	FirstName     string   `json:"fname"`
	IsMarried     bool     `json:"-"`
	IsEnrolled    bool     `json:"enrolled,omitempty"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Hours  int    `json:"coursehours"`
	Number int    `json:"coursenum"`
}

func newStudent(studentId int, lastName, middleInitial, firstName string, isMarried, isEnrolled bool) student {
	s := student{
		StudentId:     studentId,
		LastName:      lastName,
		MiddleInitial: middleInitial,
		FirstName:     firstName,
		IsMarried:     isMarried,
		IsEnrolled:    isEnrolled,
	}

	return s
}

func main() {
	s1 := newStudent(13, "Hayati", "ARMD", "Mustafa", false, true)

	student1, err := json.MarshalIndent(s1, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", student1)

	s2 := newStudent(26, "Tudi", "H", "Malena", false, false)
	c := course{Name: "world lit", Hours: 3, Number: 101}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "algebra", Hours: 8, Number: 202}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "house keeping", Hours: 1, Number: 102}
	s2.Courses = append(s2.Courses, c)

	student2, err := json.MarshalIndent(s2, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", student2)
}
