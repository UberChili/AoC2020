package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Group struct {
	answers map[rune]int
	people  int
}

func main() {
	fmt.Println("Hello there!")

	groups := GetGroups("input.txt")

	count := 0
	second_count := 0
	for i, g := range groups {
		count += CountUniqueAnswers(g)
		fmt.Printf("--- Group %d ---\n%s\n\n", i+1, g)
		fmt.Println("Unique answers:", CountUniqueAnswers(g))
		numPeople := strings.Count(g, "\n") + 1
		second_count += AnswersByEveryoneInGroup(numPeople, g)
		fmt.Println("Votes where all people answered:", AnswersByEveryoneInGroup(numPeople, g))
	}
	fmt.Println("Count of unique answers:", count)
	fmt.Println("Count of part 2 ffs Idk how tf I'm naming shit but here's the fucking number:", second_count)
}

// For Part 2, we need to track *how many times* each question was answered yes
// (i.e. count per question)
func AnswersByEveryoneInGroup(people int, text string) int {
	// groups := make([]Group, 0)
	// group := Group{}

	answers := make(map[rune]int)

	// "populating" map
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			answers[char] += 1
		}
	}

	// We need to know *how many people* voted, that's why we get people as an argument

	// Go through map and keep track of answers which
	// answer_times == number of people in group (times_required)
	count := 0
	for _, times := range answers {
		if times == people {
			count += 1
		}
	}
	return count
}

// So in Part 1 we only needed to keep track of "seen at least once"
func CountUniqueAnswers(group string) int {
	answers := make(map[rune]bool)

	for _, char := range group {
		if char >= 'a' && char <= 'z' {
			answers[char] = true
		}
	}

	return len(answers)
}

func GetGroups(filepath string) []string {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error opening file:", filepath, err)
	}

	groups := strings.Split(string(data), "\n\n")

	for i := range groups {
		groups[i] = strings.TrimSpace(groups[i])
	}

	return groups
}
