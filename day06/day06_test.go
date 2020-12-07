package main

import (
	"strings"
	"testing"
)

const testInputString = `abc

a
b
c

ab
ac

a
a
a
a

b`

func Test_day06Part1(t *testing.T) {
	in := strings.Split(testInputString, "\n\n")
	expected := 11
	actual := day06Part1(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}

func Test_day06Part2(t *testing.T) {
	in := strings.Split(testInputString, "\n\n")
	expected := 6
	actual := day06Part2(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}
