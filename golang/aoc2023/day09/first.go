package day09

import (
	"fmt"
	"syodage/aoc2023/utils"
)

func findNextVal(in []int) int {
	// fmt.Printf("in: %+v\n", in)
	allZero := true
	size := len(in)
	if size < 2 {
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
		return diff[len(diff)-1]
	}

	return diff[len(diff)-1] + findNextVal(diff)
}
func FirstPart(inputs []string) (int, error) {
	sumOfHist := 0

	for _, in := range inputs {
		values, err := utils.ToInts(utils.SplitBySpaces(in))
		if err != nil {
			return 0, err
		}

		nextVal := values[len(values)-1] + findNextVal(values)
		fmt.Printf("Sub ans: %d\n", nextVal)
		sumOfHist += nextVal
	}
	fmt.Printf("Ans: %d\n", sumOfHist)
	return sumOfHist, nil

}
