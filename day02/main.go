package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	rows := strings.Split(string(content), "\n")

	passwords := make([]passwordValidation, 0, len(rows))
	for _, row := range rows {
		password, err := getPasswordValidationFromRow(row)
		if err != nil {
			log.Fatalf("Failed to parse password validation from row: %s", err)
		}
		passwords = append(passwords, password)
	}

	day02Part1Answer := day02Part1(passwords)
	log.Printf("The answer for part 1 is: %d", day02Part1Answer)

	day02Part2Answer := day02Part2(passwords)
	log.Printf("The answer for part 2 is: %d", day02Part2Answer)
}

type passwordValidation struct {
	min      int
	max      int
	char     string
	password string
}

func getPasswordValidationFromRow(row string) (passwordValidation, error) {
	var password passwordValidation

	split1 := strings.Split(row, "-")
	min, err := strconv.Atoi(split1[0])
	if err != nil {
		return passwordValidation{}, err
	}
	password.min = min

	split2 := strings.Split(split1[1], " ")
	max, err := strconv.Atoi(split2[0])
	if err != nil {
		return passwordValidation{}, err
	}
	password.max = max
	password.password = split2[2]

	split3 := strings.Split(split2[1], ":")
	password.char = split3[0]
	return password, nil
}

func day02Part1(in []passwordValidation) int {
	var validPasswordCount int
	for _, p := range in {
		if validatePasswordPart1(p) {
			validPasswordCount++
		}
	}
	return validPasswordCount
}

func validatePasswordPart1(in passwordValidation) bool {
	var validCharCount int
	for _, passChar := range in.password {
		if string(passChar) == in.char {
			validCharCount++
		}
	}
	return in.min <= validCharCount && validCharCount <= in.max
}

func day02Part2(in []passwordValidation) int {
	var validPasswordCount int
	for _, p := range in {
		if validatePasswordPart2(p) {
			validPasswordCount++
		}
	}
	return validPasswordCount
}

func validatePasswordPart2(in passwordValidation) bool {
	firstMatch := string(in.password[in.min-1]) == in.char
	secondMatch := string(in.password[in.max-1]) == in.char
	if firstMatch != secondMatch {
		return true
	}
	return false
}
