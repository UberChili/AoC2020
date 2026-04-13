package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello there!")
}

func GetGroups(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error when opening file:", filepath, err.Error())
		return nil
	}
	defer file.Close()

	result := make([]string, 100)

	// We start a scanner and we send in a custom policy to scan/split the file
	scanner := bufio.NewScanner(file)
	scanner.Split(ScanBetweenEmptyLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

// Function to scan the input and get slices of strings separated by blank lines, which are the input
func ScanBetweenEmptyLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	// don't know how to do this
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it
	if atEOF {
		return len(data), dropCR(data), nil
	}
	return 0, nil, nil
}

func dropCR(b []byte) []byte {
	panic("unimplemented")
}
