package main

import (
	"log"
	"strconv"
	"strings"
)

// We now use a struct to keep some information abut the current bag (node)
type BagContent struct {
	color  string
	amount int
}

func get_amount_bags_inside(graph map[string][]BagContent, bag string) int {
	if graph[bag] == nil {
		return 0
	}
	count := 0
	for _, child := range graph[bag] {
		count += child.amount + (child.amount * get_amount_bags_inside(graph, child.color))
	}

	return count
}

// Parsing function for part 2
// Now we need the numbers (the amount of 'this' bag that are contained)
func parse_bag_line_part2(text string) (string, []BagContent) {
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
	child_bags := []BagContent{}
	for seq := range strings.SplitSeq(rest, ",") {
		// Remove leading whitespace that could be produced for subsequent sequences after splitting
		seq := strings.TrimSpace(seq)

		// splitting on spaces to get individual words
		bag_name_splitted := strings.Split(seq, " ")

		// Getting the amount of this bag
		amount, err := strconv.Atoi(bag_name_splitted[0])
		if err != nil {
			log.Fatal("Something went wrong getting bag amount")
		}

		// formatting the actual name/color of the bag per sequence
		bag := bag_name_splitted[1] + " " + bag_name_splitted[2]
		child_bags = append(child_bags, BagContent{bag, amount})
	}
	return bag, child_bags
}
