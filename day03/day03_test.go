package main

import (
	"strings"
	"testing"
)

var testInput = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func Test_day03Part1(t *testing.T) {
	expected := 7

	in := strings.Split(testInput, "\n")

	actual := day03Part1(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}

func Test_day03Part2(t *testing.T) {
	expected := 336

	in := strings.Split(testInput, "\n")

	actual := day03Part2(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}
