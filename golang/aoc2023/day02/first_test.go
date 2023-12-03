package day02

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

const testInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

// Game #: <set_1>; <set_2>; ...

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
			input: strings.Split(testInput, "\n"),
			want:  8,
		},
		{
			name:  "PuzzleInput",
			input: firstInput,
			want:  2061,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := First(test.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			if got != test.want {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
