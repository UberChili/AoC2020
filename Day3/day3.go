package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Right 3, down 1 (part 1)
func calc_trees(grid []string, right, down int) int {
	trees := 0
	x := 0

	for i := 0; i < len(grid); i += down {
		if grid[i][x] == '#' {
			trees += 1
		}
		x = (x + right) % len(grid[i])
	}

	return trees
}

func main() {
	lines := readLines("input.txt")

	// Part 1
	treesA := calc_trees(lines, 3, 1)
	// Right 1, Down 1
	treesB := calc_trees(lines, 1, 1)
	// Right 5, Down 1
	treesC := calc_trees(lines, 5, 1)
	// Right 7, Down 1
	treesD := calc_trees(lines, 7, 1)
	// Right 1, Down 2
	treesE := calc_trees(lines, 1, 2)

	fmt.Println("Trees:", (treesA * treesB * treesC * treesD * treesE))
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
