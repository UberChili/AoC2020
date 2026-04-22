package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	inst string
	num  int
}

func main() {
	lines := get_lines("input.txt")

	instructions := make([]Instruction, 0)
	for _, line := range lines {
		instructions = append(instructions, parse_line_into_instruction(line))
	}

	inst_indices := make(map[int]int)

	accumulator := 0
	index := 0
	for index < len(instructions) {
		fmt.Printf("index: %d | instruction: %#v | accumulator: %d\n", index, instructions[index], accumulator)
		if _, ok := inst_indices[index]; ok {
			fmt.Println("Accumulator:", accumulator)
			break
		}
		current := instructions[index]
		switch current.inst {
		case "jmp":
			inst_indices[index] = 1
			index += current.num
		case "nop":
			inst_indices[index] = 1
			index += 1
		case "acc":
			inst_indices[index] = 1
			index += 1
			accumulator += current.num
		}
	}
}

func parse_line_into_instruction(line string) Instruction {
	splitted := strings.Split(line, " ")

	value, err := strconv.Atoi(splitted[1])
	if err != nil {
		log.Fatal("Could not convert:", err)
	}
	return Instruction{inst: splitted[0], num: value}
}

func get_lines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Could not open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := []string{}

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
