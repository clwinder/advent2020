package main

import "testing"

func Test_day01(t *testing.T) {
	in := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	expected := 514579

	actual := day01(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}
