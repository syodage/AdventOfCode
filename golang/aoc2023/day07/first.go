package day07

import (
	"fmt"
	"slices"
	"strconv"
	"syodage/aoc2023/utils"
	"unicode"
)

type CCard struct {
	Cards string
	T     int // high card = 1, ..., five of a kind = 5
	Bid   int
	// Rank  int
}

func (c *CCard) String() string {
	return fmt.Sprintf("{Cards: %q, type: %d}", c.Cards, c.T)
}

func NewCCard(cards string, bid int, weekJ bool) *CCard {
	return &CCard{cards, FindType(cards, weekJ), bid}
}

func FindType(cards string, wildJ bool) int {
	mm := make(map[rune]int)
	for _, r := range cards {
		v, ok := mm[r]
		if !ok {
			mm[r] = 1
			continue
		}
		mm[r] = v + 1
	}

	if wildJ {
		rj := rune('J')
		jc, ok := mm[rj]
		if ok {
			hk := rj
			hv := 0
			for k, v := range mm {
				if k == rj {
					continue
				}
				if v > hv {
					hv = v
					hk = k
				}
			}
			delete(mm, rj)
			mm[hk] = mm[hk] + jc
		}
	}

	t := 0
	dT := len(mm)
	if dT == 1 {
		t = 7 // five of a kind
	} else if dT == 4 {
		t = 2 // one pair
	} else if dT == 5 {
		t = 1 // High card
	} else if dT == 2 {
		// four of a kind or fullhouse
		for _, v := range mm {
			if v == 4 || v == 1 {
				t = 6 // four of a kind
			} else {
				t = 5 // full house
			}
		}
	} else if dT == 3 {
		for _, v := range mm {
			if v == 3 {
				t = 4 // three of a kind
			} else if v == 2 {
				t = 3 // two pairs
			}
		}
	}

	return t
}

func FirstPart(inputs []string) (int, error) {
	cList, err := parseInput(inputs, false)
	if err != nil {
		return -1, err
	}

	SortCard(cList, false)
	// fmt.Printf("Cards:\n %+v\n", cList)
	totalWining := 0
	for i, c := range cList {
		// fmt.Printf("bid * rank = %d * %d", c.Bid, i+1)
		totalWining += (c.Bid * (i + 1))
	}
	return totalWining, nil
}

func SortCard(cList []*CCard, weekJ bool) []*CCard {
	slices.SortFunc(cList, func(a, b *CCard) int {
		if a.T != b.T {
			return a.T - b.T
		}
		// fmt.Printf("Equals : %s, %s\n", a.Cards, b.Cards)
		// we need to compare the cards from begining to find the strong cards
		for i := range a.Cards {
			an := toInt(a.Cards[i], weekJ)
			bn := toInt(b.Cards[i], weekJ)
			if an != bn {
				return an - bn
			}
		}
		return 0
	})
	return cList
}

func toInt(b byte, weekJ bool) int {
	//A, K, Q, J, T,
	if unicode.IsDigit(rune(b)) {
		n, err := strconv.Atoi(string(b))
		if err != nil {
			fmt.Printf("Error: %v", err)
			return -2
		}
		return n
	}

	if rune(b) == rune('A') {
		return 14
	}
	if rune(b) == rune('K') {
		return 13
	}
	if rune(b) == rune('Q') {
		return 12
	}
	if rune(b) == rune('J') {
		if weekJ {
			return 1
		}
		return 11
	}
	if rune(b) == rune('T') {
		return 10
	}

	return -1
}

func parseInput(inputs []string, weekJ bool) ([]*CCard, error) {
	var cards []*CCard
	for _, input := range inputs {
		sp := utils.SplitBySpaces(input)
		bid, err := strconv.Atoi(sp[1])
		if err != nil {
			return nil, err
		}
		cards = append(cards, NewCCard(sp[0], bid, weekJ))
	}
	return cards, nil
}
