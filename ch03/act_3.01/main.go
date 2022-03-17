package main

import (
	"fmt"
)

func calcSalesTax(cost float64, taxRate float64) float64 {
	return cost * (taxRate / 100)
}

func main() {
	items := map[string][2]float64{
		"Cake":   {0.99, 7.5},
		"Milk":   {2.75, 1.5},
		"Butter": {0.87, 2},
	}

	var salesTaxTotal float64

	for _, v := range items {
		salesTaxTotal += calcSalesTax(v[0], v[1])
	}

	fmt.Println("Sales tax total: ", salesTaxTotal)
}
