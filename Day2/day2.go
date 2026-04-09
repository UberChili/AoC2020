package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rules struct {
	password  string
	lower     int
	upper     int
	character string
}

func (r *Rules) GetPasswd(line string) {
	password, _ := strings.CutPrefix(line, " ")
	r.password = password
}

func (r *Rules) GetMetrics(text string) {
	splitted := strings.Split(text, " ")
	ranges := splitted[0]
	numbers := strings.Split(ranges, "-")

	character := splitted[1]

	r.lower, _ = strconv.Atoi(numbers[0])
	r.upper, _ = strconv.Atoi(numbers[1])
	r.character = character
}

func (r Rules) CheckPartTwoIsValid() bool {
	password := r.password
	lower_bound := string(password[r.lower-1])
	upper_bound := string(password[r.upper-1])

	if lower_bound == r.character || upper_bound == r.character {
		if lower_bound == upper_bound {
			return false
		} else {
			return true
		}
	}
	return false
}

func main() {
	filename := "input.txt"

	// getting lines as is
	lines := readLines(filename)
	if lines == nil {
		fmt.Printf("No lines got: ")
	}

	var correct_passwords int

	for _, line := range lines {
		// Creating an instance of line rules
		line_rules := Rules{}

		// First split
		rules, password, _ := strings.Cut(line, ":")

		// Getting password
		line_rules.GetPasswd(password)

		// getting "metrics"
		line_rules.GetMetrics(rules)

		// check if password is correct and count all correct passwords for Part 1
		// char_count := strings.Count(line_rules.password, line_rules.character)
		// if char_count <= line_rules.upper && char_count >= line_rules.lower {
		// 	correct_passwords += 1
		// }

		// Part 2
		if line_rules.CheckPartTwoIsValid() {
			correct_passwords += 1
		}

		// "Debugging" printf statement
		// fmt.Printf("%q, %q. Letter: %q. Metrics: %d-%d\n", rules, line_rules.password, line_rules.character, line_rules.lower, line_rules.upper)
	}

	fmt.Println("Correct passwords:", correct_passwords)
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
