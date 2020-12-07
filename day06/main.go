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

	answerPart1 := day06Part1(allAnswers)
	log.Printf("The answer for part 1 is %d", answerPart1)

	answerPart2 := day06Part2(allAnswers)
	log.Printf("The answer for part 2 is %d", answerPart2)
}

func day06Part1(allAnswers []string) int {
	var countYes int
	for _, a := range allAnswers {
		countYes = countYes + countGroupYesPart1(strings.Split(a, "\n"))
	}
	return countYes
}

func countGroupYesPart1(groupAnswers []string) int {
	uniqueYes := make(map[string]struct{})
	for _, answers := range groupAnswers {
		for _, a := range answers {
			uniqueYes[string(a)] = struct{}{}
		}
	}
	return len(uniqueYes)
}

func day06Part2(allAnswers []string) int {
	var countYes int
	for _, a := range allAnswers {
		countYes = countYes + countGroupYesPart2(strings.Split(a, "\n"))
	}
	return countYes
}

func countGroupYesPart2(groupAnswers []string) int {
	yes := make(map[string]int)
	for _, answers := range groupAnswers {
		for _, a := range answers {
			yes[string(a)] = yes[string(a)] + 1
		}
	}
	var allYes int
	for _, v := range yes {
		if v == len(groupAnswers) {
			allYes++
		}
	}
	return allYes
}
