package day07

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

const firstSampleTest = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

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
			want:  6440,
		},
		{
			name:  "PuzzleInput",
			input: firstInput,
			want:  250058342,
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
