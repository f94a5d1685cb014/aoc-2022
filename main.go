package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Day01()
	Day02()
}

func ReadInput(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, "opening file:", err)
	}

	defer file.Close()

	var data []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return data
}
