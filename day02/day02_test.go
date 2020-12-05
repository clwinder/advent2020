package main

import "testing"

func Test_day02Part1(t *testing.T) {
	in := []passwordValidation{
		passwordValidation{
			min:      1,
			max:      3,
			char:     "a",
			password: "abcde",
		},
		passwordValidation{
			min:      1,
			max:      3,
			char:     "b",
			password: "cdefg",
		},
		passwordValidation{
			min:      2,
			max:      9,
			char:     "c",
			password: "ccccccccc",
		},
	}
	expected := 2

	actual := day02Part1(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}

func Test_day02Part2(t *testing.T) {
	in := []passwordValidation{
		passwordValidation{
			min:      1,
			max:      3,
			char:     "a",
			password: "abcde",
		},
		passwordValidation{
			min:      1,
			max:      3,
			char:     "b",
			password: "cdefg",
		},
		passwordValidation{
			min:      2,
			max:      9,
			char:     "c",
			password: "ccccccccc",
		},
	}
	expected := 1

	actual := day02Part2(in)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}

func Test_getPasswordValidationFromRow(t *testing.T) {
	tests := map[string]struct {
		row                        string
		expectedPasswordValidation passwordValidation
	}{
		"abcde": {
			row: "1-3 a: abcde",
			expectedPasswordValidation: passwordValidation{
				min:      1,
				max:      3,
				char:     "a",
				password: "abcde",
			},
		},
		"cdefg": {
			row: "1-3 b: cdefg",
			expectedPasswordValidation: passwordValidation{
				min:      1,
				max:      3,
				char:     "b",
				password: "cdefg",
			},
		},
		"ccccccccc": {
			row: "2-9 c: ccccccccc",
			expectedPasswordValidation: passwordValidation{
				min:      2,
				max:      9,
				char:     "c",
				password: "ccccccccc",
			},
		},
		"doubleDigits": {
			row: "1-10 d: asdfg",
			expectedPasswordValidation: passwordValidation{
				min:      1,
				max:      10,
				char:     "d",
				password: "asdfg",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := getPasswordValidationFromRow(test.row)
			if err != nil {
				t.Errorf("Expected err to be nil, got %s instead", err)
			}
			if actual != test.expectedPasswordValidation {
				t.Errorf("Expected %v, got %v instead", test.expectedPasswordValidation, actual)
			}
		})
	}
}
