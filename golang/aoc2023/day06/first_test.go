package day06

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

const firstSampleTest = `Time:      7  15   30
Distance:  9  40  200`

func TestFirstPart(t *testing.T) {
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
			want:  288,
		},
		{
			name:  "PuzzleInput",
			input: firstInput,
			want:  6209190,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := FirstPart(test.input)
			if err != nil {
				t.Errorf("Returned error: %v", err)
			}

			if got != test.want {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
