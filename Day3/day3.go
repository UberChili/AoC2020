package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "example.txt"
	lines := readLines(filename)

	for _, line := range lines {
		fmt.Println(line)
	}
}

func readLines(filename string) []string {
	// your bufio shitfuckery here
	lines := make([]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file", filename)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
