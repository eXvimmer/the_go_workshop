package main

import (
	"fmt"
)

func main() {
	logLevel := "デバッグ" // a multi-byte string

	// NOTE:
	// When using range, instead of going 1 byte at a time, it moves along the
	// string 1 rune at a time. The index is the byte offset and the value is a
	// rune.

	for i, runeVal := range logLevel {
		fmt.Println(i, string(runeVal))
	}
}
