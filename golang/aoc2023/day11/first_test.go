package day11

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

const firstSampleTest = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestFirstPart(t *testing.T) {
	firstInput, err := utils.ReadLines("./first_input.txt")
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
			spaceExpandFactor: 2,
			want:              374,
		},
		{
			name:              "SampleInput",
			input:             strings.Split(firstSampleTest, "\n"),
			spaceExpandFactor: 10,
			want:              1030,
		},
		{
			name:              "SampleInput",
			input:             strings.Split(firstSampleTest, "\n"),
			spaceExpandFactor: 100,
			want:              8410,
		},
		{
			name:              "PuzzleInput",
			input:             firstInput,
			spaceExpandFactor: 2,
			want:              10885634,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := FirstPart(test.input, test.spaceExpandFactor)
			if err != nil {
				t.Errorf("Returned error: %v", err)
			}

			if got != test.want {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
