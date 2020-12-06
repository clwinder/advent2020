package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var expectedFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	passports := getPassportsFromString(string(content))

	answer := day04(passports)
	log.Printf("The answer is %d", answer)
}

type passport struct {
	fields map[string]string
}

func getPassportsFromString(in string) []passport {
	// passports separated by two new lines
	passportStrings := strings.Split(in, "\n\n")

	passports := make([]passport, 0, len(passportStrings))
	for _, p := range passportStrings {
		passportMap := make(map[string]string)
		// fields separated by spaces or new line
		fields := strings.Fields(p)
		for _, f := range fields {
			// key:value
			pair := strings.Split(f, ":")
			passportMap[pair[0]] = pair[1]
		}
		passports = append(passports, passport{fields: passportMap})
	}

	return passports
}

func day04(passports []passport) int {
	var countValidPassports int
	for _, p := range passports {
		var invalid bool
		for _, f := range expectedFields {
			if _, ok := p.fields[f]; !ok {
				invalid = true
				break
			}
		}
		if !invalid {
			countValidPassports++
		}
	}
	return countValidPassports
}
