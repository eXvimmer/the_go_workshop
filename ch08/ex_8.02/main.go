package main

import "fmt"

var budgetCategories = make(map[int]string)

func init() {
	fmt.Println("initializing our budget categories")
	budgetCategories[1] = "insurance"
	budgetCategories[2] = "mortgage"
	budgetCategories[3] = "bills"
	budgetCategories[4] = "groceries"
	budgetCategories[5] = "car payment"
	budgetCategories[6] = "vacation"
	budgetCategories[7] = "retirement"
}

func main() {
	for k, v := range budgetCategories {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
}
