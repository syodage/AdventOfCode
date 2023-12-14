package day06

import (
	"syodage/aoc2023/utils"
)

func FirstPart(inputs []string) (int, error) {
	tt, rec, err := parseInput(inputs)
	if err != nil {
		return -1, err
	}

	var beatRec []int
	for k, t := range tt {
		r := rec[k]
		beatRec = append(beatRec, numOfWaysToBeat(t, r))
	}

	m := 1
	for _, r := range beatRec {
		// fmt.Printf("r: %d\n", r)
		m *= r
	}
	return m, nil
}

func numOfWaysToBeat(t, r int) int {
	for i := 1; i < t; i++ {
		dis := i * (t - i)
		if dis > r {
			// fmt.Printf("dis: %d, beat: %d, time: %d, rec: %d\n", dis, i, t, r)
			return t - 2*i + 1
		}
	}
	return -1
}
func parseInput(inputs []string) ([]int, []int, error) {
	//Time:      7  15   30
	// Distance:  9  40  20
	tt := utils.SplitBySpaces(inputs[0])
	rr := utils.SplitBySpaces(inputs[1])

	t, err := utils.ToInts(tt[1:])
	if err != nil {
		return nil, nil, err
	}
	r, err := utils.ToInts(rr[1:])
	if err != nil {
		return nil, nil, err
	}

	return t, r, nil
}
