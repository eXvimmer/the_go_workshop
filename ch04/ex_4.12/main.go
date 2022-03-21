package main

import (
	"fmt"
)

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1    // simple copy of s1
	s3 := s1[:] // range copy of s1
	s1[3] = 99
	// NOTE: each of these slices point to the same hidden array

	return s1[3], s2[3], s3[3] // 99, 99, 99
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1           // simple copy of s1
	s1 = append(s1, 6) // append returns a new slice with more capacity
	s1[3] = 99

	return s1[3], s2[3] // 99, 4
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10) // NOTE: s1 has extra capacity
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1           // a simple copy of s1
	s1 = append(s1, 6) // there's room for 6, no new slices created
	s1[3] = 99

	return s1[3], s2[3] // 99, 99
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10) // NOTE: there's extra room for new items
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1                          // simple copy of s1
	s1 = append(s1, []int{10: 11}...) // new slice
	// we appended more values than there was available capcity, hence new slice
	s1[3] = 99

	return s1[3], s2[3] // 99, 4
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1)) // a new slice in a new address
	copied := copy(s2, s1)     // copies values only
	s1[3] = 99

	return s1[3], s2[3], copied // 99, 4, 5
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...) // a copy of values of s1 that are not linked to s1
	// s2 := append(s1[:0:0], s1...) // most memory-efficient to copy a slice
	s1[3] = 99

	return s1[3], s2[3] // 99, 4
}

func main() {
	l1, l2, l3 := linked()
	fmt.Println("Linked :", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("No Link :", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap Link :", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link: ", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)
}
