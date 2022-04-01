package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("The local current time is: ", now.Format(time.ANSIC))

	NY, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("The time in New York is: ", now.In(NY).Format(time.ANSIC))

	LA, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("The time in Los Angeles is: ", now.In(LA).Format(time.ANSIC))
}
