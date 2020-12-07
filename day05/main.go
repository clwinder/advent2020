package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const (
	numRows = 128
	numCols = 8
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	boardingPasses := strings.Split(string(content), "\n")

	answer := day05(boardingPasses)
	log.Printf("The answer is %d", answer)
}

func day05(boardingPasses []string) int {
	var maxID int
	for _, b := range boardingPasses {
		id := getSeatIDFromBoardingPass(b)
		if maxID < id {
			maxID = id
		}
	}
	return maxID
}

func getSeatIDFromBoardingPass(boardingPass string) int {
	minRow := 0
	maxRow := numRows - 1
	minCol := 0
	maxCol := numCols - 1

	for _, c := range boardingPass {
		dRow := (maxRow - minRow) / 2
		dCol := (maxCol - minCol) / 2

		switch string(c) {
		case "F":
			maxRow = maxRow - dRow - 1
		case "B":
			minRow = minRow + dRow + 1
		case "L":
			maxCol = maxCol - dCol - 1
		case "R":
			minCol = minCol + dCol + 1
		}
	}

	return 8*minRow + minCol
}
