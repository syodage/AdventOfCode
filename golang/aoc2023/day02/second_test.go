package day02

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
			input: strings.Split(testInput, "\n"),
			want:  2286,
		},
		{
			name:  "PuzzelInput",
			input: secondInput,
			want:  2286,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Second(test.input)
			if err != nil {
				t.Errorf("Error returned: %v", err)
			}
			if test.want != got {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
