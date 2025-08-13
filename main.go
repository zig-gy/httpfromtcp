package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
		return
	}

	for {
		fileBytes := make([]byte, 8)
		_, err = file.Read(fileBytes)

		if err == io.EOF {
			fmt.Println("file finished")
			os.Exit(0)
			return
		}

		if err != nil {
			fmt.Printf("error reading file: %v\n", err)
			os.Exit(1)
			return
		}

		fmt.Printf("read: %s\n", fileBytes)
	}
}
