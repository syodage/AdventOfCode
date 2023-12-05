package day03

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

func TestSecond(t *testing.T) {
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
			input: strings.Split(firstSampleTest, "\n"),
			want:  467835,
		},
		{
			name:  "PuzzleInput",
			input: secondInput,
			want:  82818007,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Second(test.input)
			if err != nil {
				t.Errorf("Returned error: %v", err)
			}

			if got != test.want {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
