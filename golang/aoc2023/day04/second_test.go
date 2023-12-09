package day04

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
			want:  30,
		},
		{
			name:  "PuzzelInput",
			input: secondInput,
			want:  13768818,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := SecondPart(test.input)
			if err != nil {
				t.Errorf("Error returned: %v", err)
			}
			if test.want != got {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
