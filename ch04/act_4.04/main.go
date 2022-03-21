package main

import "fmt"

func main() {
	dow := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	fmt.Println(dow)

	dow = append(dow[len(dow)-1:], dow[0:len(dow)-1]...)

	fmt.Println(dow)
}
