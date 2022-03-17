package main

import (
	"fmt"
)

func main() {
	/* NOTE:
	* Go stores all strings as byte collection. To be able to safely perform
	* operations with text of any kind of encoding, it should be converted from
	* a byte collection to a rune collection.
	* string -> byte collection -> rune collection
	 */

	username := "Sir_King_Über"

	// NOTE:  👇 the numbers printed out are the byte values of the string.
	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}
	fmt.Println()

	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]), " ")
	}
	fmt.Println()

	// NOTE: to safely work with interindividual characters of multi-byte string,
	// you first must convert the strings slice of byte types to a slice of rune
	// types.
	runes := []rune(username)

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
	fmt.Println()
}
