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

	answer := day01(intVals)

	log.Printf("The answer is: %d\n", answer)
}

func day01(in []int) int {
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
