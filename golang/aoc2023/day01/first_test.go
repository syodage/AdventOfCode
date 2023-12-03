package day01

import (
	"strings"
	"syodage/aoc2023/utils"
	"testing"
)

func TestFirst(t *testing.T) {
	firstInput, err := utils.ReadLines("./first_input.txt")
	if err != nil {
		t.Errorf("Error reading first input file")
	}
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "TestInput",
			input: strings.Split(testInput, "\n"),
			want:  142,
		},
		{
			name:  "RealInput",
			input: firstInput,
			want:  54634,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := First(test.input)
			if err != nil {
				t.Errorf("Day 01: returned an error: %v", err)
			}
			if got != test.want {
				t.Errorf("Got: %d, want: %d", got, test.want)
			}
		})
	}
}

const testInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
