package main

import "testing"

func Test_day01Part1(t *testing.T) {
	in := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	expected := 514579

	actual := day01Part1(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}

func Test_day01Part2(t *testing.T) {
	in := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	expected := 241861950

	actual := day01Part2(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}
