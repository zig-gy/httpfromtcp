package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(file *os.File) error {
	var curLine string

	for {
		fileBytes := make([]byte, 8)
		_, err := file.Read(fileBytes)

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return fmt.Errorf("error reading file: %v", err)

		}

		fmt.Printf("read: %s\n", fileBytes)
	}
}
