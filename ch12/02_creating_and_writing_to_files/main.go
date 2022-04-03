package main

import "os"

func main() {
	f, err := os.Create("output")
	// Create is like the touch command. If the file already exists, then it
	// truncates the file.

	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte("Using the write method.\n"))
	f.WriteString("Using the WriteString method.\n")
}
