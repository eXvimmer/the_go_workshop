package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	if len(pw) < 8 {
		return false
	}

	pwr := []rune(pw)

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwr {
		if unicode.IsUpper(v) {
			hasUpper = true
		}

		if unicode.IsLower(v) {
			hasLower = true
		}

		if unicode.IsNumber(v) {
			hasNumber = true
		}

		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("") {
		fmt.Println("Good password")
	} else {
		fmt.Println("Bad password")
	}

	if passwordChecker("m05tAf@9") {
		fmt.Println("Good password")
	} else {
		fmt.Println("Bad password")
	}

}
