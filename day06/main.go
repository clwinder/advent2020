package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	allAnswers := strings.Split(string(content), "\n\n")

	answer := day06(allAnswers)
	log.Printf("The answer is %d", answer)
}

func day06(allAnswers []string) int {
	var countYes int
	for _, a := range allAnswers {
		countYes = countYes + countGroupYes(strings.Split(a, "\n"))
	}
	return countYes
}

func countGroupYes(groupAnswers []string) int {
	uniqueYes := make(map[string]struct{})
	for _, answers := range groupAnswers {
		for _, a := range answers {
			uniqueYes[string(a)] = struct{}{}
		}
	}
	return len(uniqueYes)
}
