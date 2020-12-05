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
	treeMap := strings.Split(string(content), "\n")

	answerPart1 := day03Part1(treeMap)
	log.Printf("The answer for part 1 is %d", answerPart1)

	answerPart2 := day03Part2(treeMap)
	log.Printf("The answer for part 2 is %d", answerPart2)
}

type slope struct {
	dX, dY int
}

func day03Part1(treeMap []string) int {
	return encounteredTrees(treeMap, slope{
		dX: 3,
		dY: 1,
	})
}

func day03Part2(treeMap []string) int {
	slopes := []slope{
		slope{
			dX: 1,
			dY: 1,
		},
		slope{
			dX: 3,
			dY: 1,
		},
		slope{
			dX: 5,
			dY: 1,
		},
		slope{
			dX: 7,
			dY: 1,
		},
		slope{
			dX: 1,
			dY: 2,
		},
	}

	multiTree := 1
	for _, slope := range slopes {
		treeCount := encounteredTrees(treeMap, slope)
		multiTree = multiTree * treeCount
	}
	return multiTree
}

func encounteredTrees(treeMap []string, s slope) int {
	var i, j int
	var treeCount int
	for n := 0; n < len(treeMap); n++ {
		pos := treeMap[j][i]
		if string(pos) == "#" {
			treeCount++
		}
		i = i + s.dX
		j = j + s.dY
		if i >= len(treeMap[0]) {
			i = i - len(treeMap[0])
		}
		if j >= len(treeMap) {
			break
		}
	}
	return treeCount
}
