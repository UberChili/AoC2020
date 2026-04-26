package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := get_lines("input.txt")

	grid := copy_grid(lines)

	fmt.Println("Original grid:")
	for _, row := range grid {
		fmt.Println(string(row))
	}

	for {
		new_grid := simulation_step(grid)
		if equal_grids(grid, new_grid) {
			print_grid(new_grid)
			grid = new_grid
			break
		}
		grid = new_grid
	}

	// Finally, count taken seats
	taken_seats := count_seats(grid)
	fmt.Println("Taken seats:", taken_seats)
}

func count_seats(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			if grid[i][j] == '#' {
				count += 1
			}
		}
	}
	return count
}

func print_grid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func equal_grids(a, b [][]byte) bool {
	for i := 0; i < len(a); i += 1 {
		for j := 0; j < len(a[0]); j += 1 {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func simulation_step(grid [][]byte) [][]byte {
	bak_grid := copy_grid(grid)

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			// If a seat is empty (L) and there are no occupied seats adjacent to it,
			// the seat becomes occupied
			if grid[i][j] == 'L' {
				neighbors_taken := count_taken_neighbors(grid, i, j)
				if neighbors_taken == 0 {
					bak_grid[i][j] = '#'
				}
			}

			// If a seat is occupied (#) and four or more seats adjacent to it are also occupied,
			// the seat becomes empty.
			if grid[i][j] == '#' {
				neighbors_taken := count_taken_neighbors(grid, i, j)
				if neighbors_taken >= 4 {
					bak_grid[i][j] = 'L'
				}
			}
		}
	}
	return bak_grid
}

func count_taken_neighbors(grid [][]byte, row, col int) int {
	count := 0

	min_y := 0
	min_x := 0
	max_y := len(grid)
	max_x := len(grid[0])

	// top
	if row-1 >= min_y && grid[row-1][col] == '#' {
		count += 1
	}
	// bottom
	if row+1 < max_y && grid[row+1][col] == '#' {
		count += 1
	}
	// left
	if col-1 >= min_x && grid[row][col-1] == '#' {
		count += 1
	}
	// right
	if col+1 < max_x && grid[row][col+1] == '#' {
		count += 1
	}
	// top left
	if col-1 >= min_x && row-1 >= min_y && grid[row-1][col-1] == '#' {
		count += 1
	}
	// top right
	if row-1 >= min_y && col+1 < max_x && grid[row-1][col+1] == '#' {
		count += 1
	}
	// bottom left
	if row+1 < max_y && col-1 >= min_x && grid[row+1][col-1] == '#' {
		count += 1
	}
	// bottom right
	if row+1 < max_y && col+1 < max_x && grid[row+1][col+1] == '#' {
		count += 1
	}

	return count
}

func copy_grid(grid [][]byte) [][]byte {
	dest_grid := make([][]byte, 0, len(grid))

	for _, row := range grid {
		row_copy := make([]byte, len(row))
		copy(row_copy, row)
		dest_grid = append(dest_grid, row_copy)
	}

	return dest_grid
}

func get_lines(filename string) [][]byte {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not open file", file, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := [][]byte{}

	for scanner.Scan() {
		dest := make([]byte, len(scanner.Bytes()))
		copy(dest, scanner.Bytes())
		result = append(result, dest)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
