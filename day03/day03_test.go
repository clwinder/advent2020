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

func Test_day03(t *testing.T) {
	expected := 7

	in := strings.Split(testInput, "\n")

	actual := day03(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}
