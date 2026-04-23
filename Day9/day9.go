package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello there")

	numbers := get_numbers("input.txt")

	// for _, num := range numbers {
	// 	fmt.Println(num)
	// }

	result, _ := simulate_sliding_window(numbers, 25)
	fmt.Println(result)
}

// The offset given is 25
func simulate_sliding_window(numbers []int, offset int) (int, bool) {
	// start := offset - 1
	found := false
	for i := offset; i < len(numbers); i++ {
		window := numbers[i-offset : i]
		// fmt.Println("Window len:", len(window), ":", window)
		for j := range window {
			for k := j + 1; k < len(window); k++ {
				if numbers[i] == window[j]+window[k] {
					found = true
					break
				}
			}
		}
		fmt.Printf("i: %d numbers[i]: %d found: %b\n", i, numbers[i], found)
		if found == false {
			return numbers[i], true
		}
		found = false
	}
	return 0, false
}

func get_numbers(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Could not read file:", err)
	}
	defer file.Close()

	result := []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Could not convert number:", scanner.Text(), " ", err)
		}

		result = append(result, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
