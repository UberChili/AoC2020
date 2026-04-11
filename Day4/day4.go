package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var RequiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var ValidHCLLetters = []string{"a", "b", "c", "d", "e", "f"}

type Passport struct {
	fields map[string]string
	valid  bool
}

func (p Passport) New(text string) Passport {
	var fields_map = make(map[string]string)
	for pair := range strings.FieldsSeq(text) {
		fields := strings.Split(pair, ":")
		key := fields[0]
		value := fields[1]

		// putting inside map
		fields_map[key] = value
	}

	valid := true
	passport := Passport{fields: fields_map, valid: valid}

	// Checking if valid
	// So, for part 2 is not enough to check if all fields (except for that one) are present
	// We need to add some rules
	for _, field := range RequiredFields {
		got, ok := fields_map[field]
		if !ok {
			passport.valid = false
			break
		}
		switch got {
		case "byr":
			if len(got) != 4 {
				passport.valid = false
				return passport
			}
			year, err := strconv.Atoi(got)
			if err != nil {
				fmt.Println("Error with year:", year, err.Error())
				valid = false
				return passport
			}
			if year < 1920 || year > 2002 {
				valid = false
			}
		case "iyr":
			if len(got) != 4 {
				passport.valid = false
				return passport
			}
			year, err := strconv.Atoi(got)
			if err != nil {
				fmt.Println("Error with year:", year, err.Error())
				valid = false
				return passport
			}
			if year < 2010 || year > 2020 {
				valid = false
			}
		case "eyr":
			if len(got) != 4 {
				passport.valid = false
				return passport
			}
			year, err := strconv.Atoi(got)
			if err != nil {
				fmt.Println("Error with year:", year, err.Error())
				valid = false
				return passport
			}
			if year < 2020 || year > 2030 {
				valid = false
				return passport
			}
		case "hgt":
			if !strings.Contains(got, "cm") && !strings.Contains(got, "in") {
				valid = false
				return passport
			}
			if strings.Contains(got, "cm") {
				height, err := strconv.Atoi(strings.TrimSuffix(got, "cm"))
				if err != nil {
					fmt.Println("Error with height:", got, err.Error())
					valid = false
					return passport
				}
				if height < 150 || height > 193 {
					valid = false
					return passport
				}
			}
			if strings.Contains(got, "in") {
				height, err := strconv.Atoi(strings.TrimSuffix(got, "in"))
				if err != nil {
					fmt.Println("Error with height:", got, err.Error())
					valid = false
					return passport
				}
				if height < 59 || height > 76 {
					valid = false
					return passport
				}
			}
		case "hcl":
			if !strings.Contains(got, "#") {
				valid = false
				return passport
			}
			characters := strings.TrimPrefix(got, "#")
			if len(characters) != 6 {
				valid = false
				return passport
			}
			for _, c := range characters {

			}
		}
	}

	return passport
}

func main() {
	batches := read_passports("input.txt")

	passports_arr := make([]Passport, 0)
	for _, batch := range batches {
		passport := Passport{}
		passport = passport.New(batch)
		passports_arr = append(passports_arr, passport)
	}

	count := 0
	for _, pass := range passports_arr {
		if pass.valid == true {
			count += 1
		}
	}

	fmt.Println("Valid passports:", count)
}

func read_passports(filename string) []string {
	if filename == "" {
		fmt.Println("Error in filename. Can't be empty")
		return nil
	}

	// passports := make([]Passport, 0)
	passports := make([]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error when opening file:", err.Error())
	}
	defer file.Close()

	// Here we send in a custom policy to scan/split the file
	scanner := bufio.NewScanner(file)
	scanner.Split(ScanBetweenEmptyLines)

	for scanner.Scan() {
		passports = append(passports, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return passports
}

func ScanBetweenEmptyLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	// don't know how to do this
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it
	if atEOF {
		return len(data), dropCR(data), nil
	}
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
