package main

import "testing"

func Test_getSeatIDFromBoardingPass(t *testing.T) {
	tests := map[string]struct {
		boardingPass string
		expectedID   int
	}{
		"FBFBBFFRLR": {
			boardingPass: "FBFBBFFRLR",
			expectedID:   357,
		},
		"BFFFBBFRRR": {
			boardingPass: "BFFFBBFRRR",
			expectedID:   567,
		},
		"FFFBBBFRRR": {
			boardingPass: "FFFBBBFRRR",
			expectedID:   119,
		},
		"BBFFBBFRLL": {
			boardingPass: "BBFFBBFRLL",
			expectedID:   820,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getSeatIDFromBoardingPass(test.boardingPass)
			if actual != test.expectedID {
				t.Errorf("Expected %d, got %d instead", test.expectedID, actual)
			}
		})
	}
}
