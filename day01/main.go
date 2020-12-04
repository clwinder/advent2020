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
	values := strings.Split(string(content), "\n")

	var intVals []int
	for _, v := range values {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %s", err)
		}
		intVals = append(intVals, intVal)
	}

	answerPart1 := day01Part1(intVals)
	log.Printf("The answer for part 1 is: %d\n", answerPart1)

	answerPart2 := day01Part2(intVals)
	log.Printf("The answer for part 2 is: %d\n", answerPart2)
}

func day01Part1(in []int) int {
	for i := range in {
		for j := 0; j < len(in); j++ {
			if i == j {
				continue
			}
			if in[i]+in[j] == 2020 {
				return in[i] * in[j]
			}
		}
	}
	return 0
}

func day01Part2(in []int) int {
	for i := range in {
		for j := 0; j < len(in); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(in); k++ {
				if i == k {
					continue
				}
				if j == k {
					continue
				}
				if in[i]+in[j]+in[k] == 2020 {
					return in[i] * in[j] * in[k]
				}
			}
		}
	}
	return 0
}
