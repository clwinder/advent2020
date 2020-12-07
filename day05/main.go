package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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

	answerPart1 := day05Part1(boardingPasses)
	log.Printf("The answer for part 1 is %d", answerPart1)

	answerPart2, err := day05Part2(boardingPasses)
	if err != nil {
		log.Fatalf("Failed to find answer to part 2: %s", err)
	}
	log.Printf("The answer for part 2 is %d", answerPart2)
}

func day05Part1(boardingPasses []string) int {
	var maxID int
	for _, b := range boardingPasses {
		id := getSeatIDFromBoardingPass(b)
		if maxID < id {
			maxID = id
		}
	}
	return maxID
}

func day05Part2(boardingPasses []string) (int, error) {
	seatIDs := make([]int, 0, len(boardingPasses))
	for _, b := range boardingPasses {
		seatIDs = append(seatIDs, getSeatIDFromBoardingPass(b))
	}
	sort.Ints(seatIDs)

	var mySeatIDs []int
	for i := 1; i < len(seatIDs)-1; i++ {
		if seatIDs[i-1] == seatIDs[i]-2 {
			mySeatIDs = append(mySeatIDs, seatIDs[i]-1)
		}
	}
	if len(mySeatIDs) != 1 {
		return 0, fmt.Errorf("failed to find my seat, got %d options", len(mySeatIDs))
	}
	return mySeatIDs[0], nil
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
