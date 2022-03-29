package main

import (
	"fmt"
)

var budgetCategories = make(map[int]string)
var payeeToCategory = make(map[string]int)

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

func init() {
	fmt.Println("Assign our Payees to categories")
	payeeToCategory["Nationwide"] = 1
	payeeToCategory["BBT Loan"] = 2
	payeeToCategory["First Energy Electric"] = 3
	payeeToCategory["Ameriprise Financial"] = 4
	payeeToCategory["Walt Disney World"] = 5
	payeeToCategory["ALDI"] = 7
	payeeToCategory["Martins"] = 7
	payeeToCategory["Wal Mart"] = 7
	payeeToCategory["Chevy Loan"] = 8
}

func main() {
	for k, v := range budgetCategories {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}

	for k, v := range payeeToCategory {
		fmt.Printf("Payee: %s, Category: %s\n", k, budgetCategories[v])
	}
}
