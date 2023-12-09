package day04

import (
	"math"
	"strings"
	"syodage/aoc2023/utils"
)

type Card struct {
	ID      string
	WinNums []string
	Nums    []string
}

func (c *Card) MatchNums() []string {
	var res []string
	for _, w := range c.WinNums {
		for _, n := range c.Nums {
			if w == n {
				res = append(res, w)
			}
		}
	}
	return res
}

func FirstPart(inputs []string) (int, error) {
	cards := parseInputs(inputs)

	sumOfPoints := 0
	for _, card := range cards {
		matchCount := len(card.MatchNums())
		if matchCount > 0 {
			points := math.Pow(2, float64(matchCount)-1)
			sumOfPoints += int(points)
		}
	}
	return sumOfPoints, nil
}

func parseInputs(inputs []string) []*Card {
	cards := []*Card{}
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	for _, input := range inputs {
		// part 1: Card 1
		// part 2:  41 48 83 86 17 | 83 86  6 31 17  9 48 53
		cs := strings.Split(input, ":")
		id := strings.FieldsFunc(cs[0], utils.RuneSep)[1]

		nSp := strings.Split(cs[1], "|")
		wSp := strings.FieldsFunc(nSp[0], utils.RuneSep)
		ySp := strings.FieldsFunc(nSp[1], utils.RuneSep)
		cards = append(cards, &Card{id, wSp, ySp})
	}
	return cards
}
