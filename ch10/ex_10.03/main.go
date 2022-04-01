package main

import (
	"fmt"
	"time"
)

func timeDiff(timezone string) (string, string) {
	current := time.Now()
	remoteZone, _ := time.LoadLocation(timezone)
	remoteTime := current.In(remoteZone)

	fmt.Println("The current time is: ", current.Format(time.ANSIC))
	fmt.Println("The timezone: ", timezone, "time is: ", remoteTime)

	return current.Format(time.ANSIC), remoteTime.Format(time.ANSIC)
}

func main() {
	fmt.Println(timeDiff("Asia/Tehran"))
	fmt.Println(timeDiff("America/New_York"))
}
