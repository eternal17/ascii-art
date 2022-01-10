package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// opening the file in read-only mode. The file must exist (in the current working directory)
	file, _ := os.Open("standard.txt")

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanned := bufio.NewScanner(file)

	scanned.Split(bufio.ScanLines)

	var lines []string

	for scanned.Scan() {
		lines = append(lines, scanned.Text())
	}

	file.Close()

	asciiChrs := make(map[int][]string)

	dec := 31

	for _, line := range lines {
		if line == "" {
			dec++
		} else {
			asciiChrs[dec] = append(asciiChrs[dec], line)
		}
	}
	// string that user inputs
	userInput := os.Args[1]
	Newline(userInput, asciiChrs)
}

// Newline function prints string horizontally and with new line if user specifies
func Newline(n string, y map[int][]string) {
	replaceNewline := strings.ReplaceAll(n, "\\n", "\n")
	wordsSlice := strings.Split(replaceNewline, "\n")
	for _, word := range wordsSlice {
		for j := 0; j < len(y[32]); j++ {
			for _, letter := range word {
				fmt.Print(y[int(letter)][j])
			}
			fmt.Println()
		}
	}
}
