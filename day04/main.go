package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
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

	answerPart1 := day04Part1(passports)
	log.Printf("The answer for part 1 is %d", answerPart1)

	answerPart2 := day04Part2(passports)
	log.Printf("The answer for part 2 is %d", answerPart2)
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

func day04Part1(passports []passport) int {
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

func day04Part2(passports []passport) int {
	var countValidPassports int
	for _, p := range passports {
		var invalid bool
		for _, f := range expectedFields {
			v, ok := p.fields[f]
			if !ok {
				invalid = true
				break
			}
			switch f {
			case "byr":
				if !validateByr(v) {
					invalid = true
					break
				}
			case "iyr":
				if !validateIyr(v) {
					invalid = true
					break
				}
			case "eyr":
				if !validateEyr(v) {
					invalid = true
					break
				}
			case "hgt":
				if !validateHgt(v) {
					invalid = true
					break
				}
			case "hcl":
				if !validateHcl(v) {
					invalid = true
					break
				}
			case "ecl":
				if !validateEcl(v) {
					invalid = true
					break
				}
			case "pid":
				if !validatePid(v) {
					invalid = true
					break
				}
			}
		}
		if !invalid {
			countValidPassports++
		}
	}
	return countValidPassports
}

func validateByr(byr string) bool {
	// four digits; at least 1920 and at most 2002.
	n, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}

	return 1920 <= n && n <= 2002
}

func validateIyr(iyr string) bool {
	// four digits; at least 2010 and at most 2020.
	n, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}

	return 2010 <= n && n <= 2020
}

func validateEyr(eyr string) bool {
	// four digits; at least 2020 and at most 2030.
	n, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}

	return 2020 <= n && n <= 2030
}

func validateHgt(hgt string) bool {
	// a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	if strings.Contains(hgt, "cm") {
		h, err := strconv.Atoi(strings.Split(hgt, "cm")[0])
		if err != nil {
			return false
		}
		return 150 <= h && h <= 193
	}
	if strings.Contains(hgt, "in") {
		h, err := strconv.Atoi(strings.Split(hgt, "in")[0])
		if err != nil {
			return false
		}
		return 59 <= h && h <= 76
	}
	return false
}

func validateHcl(hcl string) bool {
	// a # followed by exactly six characters 0-9 or a-f.
	ok, err := regexp.MatchString(`#[a-f\d]{6}`, hcl)
	if err != nil {
		log.Printf("error matching regex: %s\n", err)
		return false
	}
	return ok
}

func validateEcl(ecl string) bool {
	// exactly one of: amb blu brn gry grn hzl oth.
	return ecl == "amb" ||
		ecl == "blu" ||
		ecl == "brn" ||
		ecl == "gry" ||
		ecl == "grn" ||
		ecl == "hzl" ||
		ecl == "oth"
}

func validatePid(pid string) bool {
	// a nine-digit number, including leading zeroes.
	if len(pid) != 9 {
		return false
	}
	_, err := strconv.Atoi(pid)
	return err == nil
}
