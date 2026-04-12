package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	fmt.Println("Hello there!")

	// example := "FBFBBFFRLR"
	seats := read_lines("input.txt")

	// Part 1 requires us to calculate seat ID's and then
	// just find the highest seat ID:
	seats_ids := make([]int, 0)
	for _, seat := range seats {
		seats_ids = append(seats_ids, sliceToDecimal(parseSeat(seat)))
	}
	highest := slices.Max(seats_ids)
	fmt.Println("Highest seat ID:", highest)

	// Part 2 requires us to find the missing Seat ID in the entire sequence
	// So we need the sorted list of Seat ID's first
	slices.Sort(seats_ids)
	for i := 1; i < len(seats_ids); i++ {
		if (seats_ids[i] - seats_ids[i-1]) != 1 {
			fmt.Println("Part 2 Seat:", seats_ids[i-1]+1)
			break
		}
	}
}

// Copied this, tbh
func sliceToDecimal(bits []int) int {
	decimal := 0
	for _, bit := range bits {
		// Shift existing total left and add the new bit
		decimal = (decimal << 1) | bit
	}
	return decimal
}

// Previously called binary_search, but we aren't actually searching nor are we actually using
// Binary search, so we'll rename it to something more descriptive
func parseSeat(seat string) []int {
	if seat == "" {
		fmt.Println("Error with seat. Can't be empty:", seat)
		return nil
	}

	binary_rows := make([]int, 0)
	// First seven characters
	for _, inst := range seat[0:7] {
		switch inst {
		case 'F':
			binary_rows = append(binary_rows, 0)
		case 'B':
			binary_rows = append(binary_rows, 1)
		}
	}

	binary_cols := make([]int, 0)

	// Last three chracters
	for _, inst := range seat[7:] {
		switch inst {
		case 'L':
			binary_cols = append(binary_cols, 0)
		case 'R':
			binary_cols = append(binary_cols, 1)
		}
	}

	result := append(binary_rows, binary_cols...)

	return result
}

func read_lines(filepath string) []string {
	if filepath == "" {
		fmt.Println("Error opening file: filename is empty")
		return nil
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", filepath, err.Error())
	}

	scanner := bufio.NewScanner(file)

	result := make([]string, 0)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
