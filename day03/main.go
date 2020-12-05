package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const (
	dX = 3
	dY = 1
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	treeMap := strings.Split(string(content), "\n")

	answer := day03(treeMap)
	log.Printf("The answer is %d", answer)
}

func day03(treeMap []string) int {
	var i, j int
	var treeCount int
	for n := 0; n < len(treeMap); n++ {
		pos := treeMap[j][i]
		if string(pos) == "#" {
			treeCount++
		}
		i = i + dX
		j = j + dY
		if i >= len(treeMap[0]) {
			i = i - len(treeMap[0])
		}
		if j >= len(treeMap) {
			break
		}
	}
	return treeCount
}
