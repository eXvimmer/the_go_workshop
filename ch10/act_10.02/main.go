package main

import (
	"fmt"
	"time"
)

func prettyDate(year, month, day, hour, minute, second, nsec int) string {
	date := time.Date(year, time.Month(month), day, hour, minute, second, nsec,
		time.UTC)

	return fmt.Sprintf("%02d:%02d:%02d %4d/%02d/%02d", date.Hour(),
		date.Minute(), date.Second(), date.Year(), date.Month(), date.Day())
}

func main() {
	fmt.Println(prettyDate(1992, 9, 26, 13, 9, 4, 2309))
	fmt.Println(prettyDate(2022, 4, 1, 16, 20, 4, 2332))
}
