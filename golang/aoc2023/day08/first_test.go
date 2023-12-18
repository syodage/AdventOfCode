package day08

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

const firstSampleTest = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const secondSampleTest = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

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
			want:  2,
		},
		{
			name:  "SampleInput2",
			input: strings.Split(secondSampleTest, "\n"),
			want:  6,
		},
		{
			name:  "PuzzleInput",
			input: firstInput,
			want:  11911,
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
