package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
		return
	}

	if err := readFile(file); err != nil {
		fmt.Printf("error reading file: %v\n", err)
		os.Exit(1)
		return
	}
}
