package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.ANSIC))

	duration := time.Duration((6*3600 + 6*60 + 6) * time.Second)
	end := now.Add(duration)

	fmt.Println("6 hours, 6 minutes and 6 seconds from now will be: ",
		end.Format(time.ANSIC))
}
