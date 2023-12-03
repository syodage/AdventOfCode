package day01

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Second(inputs []string) (int, error) {
	totalVal := 0
	for r, in := range inputs {
		res, err := readValueV2(in)
		if err != nil {
			return 0, fmt.Errorf("Line %d, Error parsing line: %v", r, err)
		}
		totalVal += res
	}
	return totalVal, nil
}

var strDigitMap = map[string]string{
	"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
}

func readValueV2(input string) (int, error) {
	currVal := ""
	currI := 0
	for currI < len(input) {
		// fmt.Printf("currI: %d, currVal: %v\n", currI, currVal)
		c := input[currI]
		rc := rune(c)
		if unicode.IsDigit(rc) {
			currVal = currVal + string(c)
		}
		// if right hand side len grater than or equals to spellout string we should check for spellout string
		// if match advance the currI to next index pass the spellout string.
		for k, v := range strDigitMap {
			if strings.HasPrefix(input[currI:], k) {
				currVal = currVal + v
			}
		}
		currI += 1
		// fmt.Printf("currI: %d, currVal: %v\n", currI, currVal)
	}

	if len(currVal) == 0 {
		return 0, fmt.Errorf("Empty curr value.")
	}

	if len(currVal) == 1 {
		currVal = currVal + currVal
	}
	numVal := string(currVal[0]) + string(currVal[len(currVal)-1])
	val, err := strconv.Atoi(string(numVal))
	if err != nil {
		return 0, fmt.Errorf("Error converting num val: %v to int", string(numVal))
	}
	return val, nil
}
