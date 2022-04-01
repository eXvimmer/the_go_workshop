package main

import (
	"fmt"
	"time"
)

// returns current date in the HH:MM:SS DD/MM/YYYY format
func prettyNow() string {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	return fmt.Sprintf("%2d:%2d:%2d %d/%d/%d", hour, minute, second, day, month, year)
}

func main() {
	fmt.Println(prettyNow())
	time.Sleep(2 * time.Second)
	fmt.Println(prettyNow())
	time.Sleep(2 * time.Second)
	fmt.Println(prettyNow())
}
