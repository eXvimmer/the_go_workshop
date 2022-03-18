package main

import "fmt"

func compArrays() (bool, bool, [10]int) {
	var arr1 [10]int
	arr2 := [...]int{9: 0} // 9 is key and 0 is the value
	arr3 := [10]int{1, 9: 10, 4: 5}

	return arr1 == arr2, arr1 == arr3, arr3
}

func main() {
	/* NOTE:
	 * Go allows you to pick the key you want for your data if you want using
	 * [<size>]<type>{<key1>:<value1>,…<keyN>:<valueN>}. Go is flexible and lets
	 * you set the keys with gaps and in any order. This ability to set values
	 * with a key is helpful if you've defined an array where the numeric keys
	 * have a specific meaning and you want to set a value for a specific key but
	 * don't need to set any of the other values.
	 */
	comp1, comp2, arr3 := compArrays()
	fmt.Println("[10]int == [...]{9:0} :", comp1)
	fmt.Println("[10]int == [10]int{1, 9: 10, 4: 5}}:", comp2)
	fmt.Println("arr3 :", arr3)
}
