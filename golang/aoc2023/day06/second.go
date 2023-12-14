package day06

import (
	"strconv"
	"strings"
	"syodage/aoc2023/utils"
)

func SecondPart(inputs []string) (int, error) {
	tt := utils.SplitBySpaces(inputs[0])
	rr := utils.SplitBySpaces(inputs[1])

	ttt := strings.Join(tt[1:], "")
	t, err := strconv.Atoi(ttt)
	if err != nil {
		return -1, err
	}

	rrr := strings.Join(rr[1:], "")
	r, err := strconv.Atoi(rrr)
	if err != nil {
		return -1, err
	}

	return numOfWaysToBeat(t, r), nil
}
