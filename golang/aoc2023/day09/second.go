package day09

import (
	"fmt"
	"syodage/aoc2023/utils"
)

func findHist(in []int) int {
	// fmt.Printf("in: %+v\n", in)
	allZero := true
	size := len(in)
	if size < 2 {
		fmt.Printf("Receive only one val: %+v", in)
		return in[0]
	}

	var diff []int
	for i, z := range in {
		if i == 0 {
			continue
		}
		d := z - in[i-1]
		if d != 0 {
			allZero = false
		}
		diff = append(diff, d)
	}

	if allZero {
		return diff[0]
	}

	return diff[0] - findHist(diff)
}
func SecondPart(inputs []string) (int, error) {

	sumOfHist := 0
	for _, in := range inputs {
		values, err := utils.ToInts(utils.SplitBySpaces(in))
		if err != nil {
			return 0, err
		}

		hist := values[0] - findHist(values)
		fmt.Printf("Sub ans: %d\n", hist)
		sumOfHist += hist
	}
	fmt.Printf("Ans: %d\n", sumOfHist)
	return sumOfHist, nil
}
