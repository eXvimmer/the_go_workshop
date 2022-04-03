package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	n := flag.String("name", "", "your first name")
	i := flag.Int("age", -1, "your age")
	m := flag.Bool("married", false, "are you married?")
	flag.Parse()

	if *n == "" { // if it has the default value (it's unset)
		fmt.Println("name is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Name:", *n)
	fmt.Println("Age: ", *i)
	fmt.Println("Married: ", *m)
}
