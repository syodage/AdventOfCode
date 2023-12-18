package day09

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

const firstSampleTest = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

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
			want:  114,
		},
		{
			name:  "PuzzleInput",
			input: firstInput,
			want:  100,
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
