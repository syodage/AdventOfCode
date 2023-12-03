package day01

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func First(inputs []string) (int, error) {
	calValues := []int{}
	for r, v := range inputs {
		val, err := readValue(v)
		// fmt.Printf("line: %d, val: %d\n", r, val)
		if err != nil {
			return 0, fmt.Errorf("Error parsing line: %d, Error: %v", r, err)
		}
		calValues = append(calValues, val)
	}

	totalCalValues := 0
	for _, v := range calValues {
		totalCalValues += v
	}
	return totalCalValues, nil
}

func readValue(line string) (int, error) {
	currVal := []rune{}
	if len(line) == 0 {
		return 0, errors.New("Empty input string")
	}
	for _, c := range line {
		if unicode.IsDigit(c) {
			if len(currVal) == 2 {
				currVal = currVal[:1]
			}
			currVal = append(currVal, c)
		}
	}
	if len(currVal) == 0 {
		return 0, errors.New("No digits found")
	}

	// if one vlaue then duplicate it
	if len(currVal) == 1 {
		currVal = append(currVal, currVal[0])
	}

	val, err := strconv.Atoi(string(currVal))
	if err != nil {
		return 0, fmt.Errorf("Error converting %+v to int", currVal)
	}

	return val, nil
}
