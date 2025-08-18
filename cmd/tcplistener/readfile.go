package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)

		var curLine string
		for {
			fileBytes := make([]byte, 8)
			_, err := f.Read(fileBytes)
			if errors.Is(err, io.EOF) {
				if curLine != "" {
					lines <- curLine
				}
				break
			}
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			parts := strings.Split(string(fileBytes), "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", curLine, parts[i])
				curLine = ""
			}
			curLine += parts[len(parts)-1]
		}
	}()
	return lines
}
