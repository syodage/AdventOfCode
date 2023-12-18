package day08

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

const thirdSampleTest = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestSecondPart(t *testing.T) {

	secondInput, err := utils.ReadLines("./second_input.txt")
	if err != nil {
		t.Errorf("Error reading input file")
	}

	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "SampleInput",
			input: strings.Split(thirdSampleTest, "\n"),
			want:  6,
		},
		{
			name:  "PuzzleInput",
			input: secondInput,
			want:  200,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := SecondPart(test.input)
			if err != nil {
				t.Errorf("Returned error: %v", err)
			}

			if got != test.want {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
