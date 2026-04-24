
package main

import "core:slice"
import "core:strconv"
import "core:strings"
import "core:log"
import "core:os"
import "core:fmt"

main :: proc() {
	numbers := get_numbers("input.txt")
	defer delete(numbers)

	slice.sort_by(numbers[:], int_order)

	for number in numbers {
		fmt.println(number)
	}
	fmt.println("Multiplication result:", traverse_numbers(numbers[:]))
}

traverse_numbers :: proc(numbers: []int) -> int {
	current := 0
	count_one := 0
	count_three := 0

	for num in numbers {
		difference := num - current
		if difference == 1 {
			count_one += 1
			current = num
			continue
		}
		if difference == 3 {
			count_three += 1
			current = num
			continue
		}
		else {
			current = num
		}
	}
	// last (my device's built-in adapter, which is always 3 higher than the highest adapter)
	count_three += 1

	return count_one * count_three
}

get_numbers ::proc(filename: string) -> [dynamic]int {
	// Skipping checking if file exists and all that for brevity
	data, err := os.read_entire_file(filename,  context.temp_allocator)
	if err != nil {
		log.fatalf("Could not read file:", err)
	}

	result := [dynamic]int{}
	it := string(data)
	for line in strings.split_lines_iterator(&it) {
		number, ok := strconv.parse_int(line)
		if !ok {
			log.fatal("Error converting to int:", number)
		}
		append(&result, number)
	}
	return result
}

int_order :: proc(lhs, rhs: int) -> bool {
	return lhs < rhs
}