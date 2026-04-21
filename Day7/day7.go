package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := get_lines("example.txt")

	// Our map (graph)
	// For part 1, we use a simple map
	// graph := make(map[string][]string)
	// For part 2, we use:
	graph := make(map[string][]BagContent)

	// Part 1. We "generate" the map:
	// for _, line := range lines {
	// 	parent, children := parse_bag_line(line)
	// 	graph[parent] = children
	// }

	// Part 2. We "generate" the map:
	for _, line := range lines {
		parent, children := parse_bag_line_part2(line)
		graph[parent] = children
	}

	// Just a printing loop to test if the map has correctly been formed
	for key, values := range graph {
		fmt.Printf("%s -> ", key)
		for _, v := range values {
			fmt.Printf("%d %s | ", v.amount, v.color)
		}
		fmt.Println()
	}

	// Loop through all bags checking if they contain shiny gold
	// This is the solution for part 1
	// bags_count := 0
	// for key := range graph {
	// 	if contains_gold(graph, key) {
	// 		bags_count += 1
	// 	}
	// }

	// fmt.Println("Bags that contain Shiny Gold:", bags_count)
}

// For part 1, we just need to go deep on every node/bag to see if it eventually contains Shiny Gold
func contains_gold(graph map[string][]string, bag string) bool {
	for _, child := range graph[bag] {
		if child == "shiny gold" {
			return true
		}
		if contains_gold(graph, child) {
			return true
		}
	}

	return false
}

// Parsing function for part 1
// We pretty much ignore all numbers and get the bag name (node) and the bags it contains (its children)
func parse_bag_line(text string) (string, []string) {
	bag, rest, found := strings.Cut(text, " bags contain ")
	if found == false {
		return "", nil
	}

	// Extracting leave bags
	//
	// If line contains the word "no", then the bag contains no number of other bags
	if strings.Contains(rest, "no") {
		return bag, nil
	}

	// We split by commas.
	// In the example, we have one or at most, two bags inside any bag.
	// But in the input, we have any number of bags
	child_bags := []string{}
	for seq := range strings.SplitSeq(rest, ",") {
		// Remove leading whitespace that could be produced for subsequent sequences after splitting
		seq := strings.TrimSpace(seq)

		//splitting on spaces to get individual words
		bag_name_splitted := strings.Split(seq, " ")

		// forming the actual name of the bag per sequence
		bag := bag_name_splitted[1] + " " + bag_name_splitted[2]
		child_bags = append(child_bags, bag)
	}
	return bag, child_bags
}

// Just the get lines function, this doesn't change for any part
func get_lines(filepath string) []string {
	// omitting file checking for brevity
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	result := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
