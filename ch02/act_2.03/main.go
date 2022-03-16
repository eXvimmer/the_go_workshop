package main

import (
	"fmt"
)

func bubbleSort(s []int) {
	l := len(s)

	for l > 0 {
		for i := 0; i < l; i++ {
			if i+1 < l && s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
		l--
	}
}

func main() {
	nums := []int{9, 4, 1, 0, 99, 13, 8, 7, 11, -20, 23, 37, 42, 47, 1992, 1371,
		2022, 2019, 2020}

	fmt.Println(nums)
	bubbleSort(nums)
	fmt.Println(nums)
}
