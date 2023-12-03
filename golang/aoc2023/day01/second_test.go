package day01

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

const testInputSecond = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

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
			name:  "Test",
			input: []string{"two1nine"},
			want:  29,
		},
		{
			name:  "Test",
			input: []string{"eightwothree"},
			want:  83,
		},
		{
			name:  "Test",
			input: []string{"abcone2threexyz"},
			want:  13,
		},
		{
			name:  "Test",
			input: []string{"xtwone3four"},
			want:  24,
		},
		{
			name:  "Test",
			input: []string{"4nineeightseven2"},
			want:  42,
		},
		{
			name:  "Test",
			input: []string{"zoneight234"},
			want:  14,
		},
		{
			name:  "Test",
			input: []string{"7pqrstsixteen"},
			want:  76,
		},
		{
			name:  "TestInupt",
			input: strings.Split(testInputSecond, "\n"),
			want:  281,
		},
		{
			name:  "RealInupt",
			input: secondInput,
			want:  53855,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Second(test.input)
			if err != nil {
				t.Errorf("Got error: %v", err)
			}

			if test.want != got {
				t.Errorf("Want: %d, got: %d", test.want, got)
			}
		})
	}
}
