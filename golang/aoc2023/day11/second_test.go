package day11

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

func TestSecondPart(t *testing.T) {

	secondInput, err := utils.ReadLines("./second_input.txt")
	if err != nil {
		t.Errorf("Error reading input file")
	}

	tests := []struct {
		name              string
		input             []string
		spaceExpandFactor int
		want              int
	}{
		{
			name:              "SampleInput",
			input:             strings.Split(firstSampleTest, "\n"),
			spaceExpandFactor: 1000000,
			want:              82000210,
		},
		{
			name:              "PuzzleInput",
			input:             secondInput,
			spaceExpandFactor: 1000000,
			want:              707505470642,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := SecondPart(test.input, test.spaceExpandFactor)
			if err != nil {
				t.Errorf("Returned error: %v", err)
			}

			if got != test.want {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
