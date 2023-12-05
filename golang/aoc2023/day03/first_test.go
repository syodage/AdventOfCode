package day03

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

const firstSampleTest = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestFirst(t *testing.T) {
	firstInput, err := utils.ReadLines("./first_input.txt")
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
			want:  4361,
		},
		{
			name:  "PuzzleInput",
			input: firstInput,
			want:  539637,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := First(test.input)
			if err != nil {
				t.Errorf("Returned error: %v", err)
			}

			if got != test.want {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
