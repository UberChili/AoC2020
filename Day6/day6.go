package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Group struct {
	answers map[rune]int
}

func main() {
	fmt.Println("Hello there!")

	groups := GetGroups("example.txt")

	count := 0
	for i, g := range groups {
		count += CountUniqueAnswers(g)
		fmt.Printf("--- Group %d ---\n%s\n\n", i+1, g)
		fmt.Println("Unique answers:", CountUniqueAnswers(g))
	}
	fmt.Println("Count of unique answers:", count)
}

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
