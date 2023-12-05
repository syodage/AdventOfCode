package day03

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Dir struct {
	dx int
	dy int
}

// (-1, -1)(0, -1)(1, -1)
// (-1, 0)(0, 0)(1, 0)
// (-1, 1)(0, 1)(1, 1)
var dirs = []Dir{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

const dot = rune('.')

func First(inputs []string) (int, error) {
	es := parseInput(inputs)
	nums := []string{}
	num := strings.Builder{}
	for r, line := range inputs {
		adjSym := false
		for c, sym := range line {
			if !unicode.IsDigit(sym) {
				if num.Len() != 0 {
					if adjSym {
						nums = append(nums, num.String())
					}
					num.Reset()
				}
				adjSym = false
				continue
			}

			if unicode.IsDigit(sym) {
				adjSym = adjSym || adjacentToASymbol(es, c, r)
				num.WriteRune(sym)
			}
		}
		// if a valid number with adjacent to a Symbol
		if num.Len() != 0 && adjSym {
			nums = append(nums, num.String())
		}
	}
	// convert strings to numbers and calculate the sum
	// fmt.Printf("Found %d valid numbers\n", len(nums))
	sum := 0
	for _, val := range nums {
		n, err := strconv.Atoi(val)
		if err != nil {
			return 0, fmt.Errorf("Error converting %q to int, err: %v", val, err)
		}
		sum += n
	}
	return sum, nil
}

func parseInput(inputs []string) [][]rune {
	es := [][]rune{}
	for _, line := range inputs {
		esr := []rune{}
		for _, r := range line {
			esr = append(esr, r)
		}
		es = append(es, esr)
	}
	return es
}

func adjacentToASymbol(es [][]rune, c, r int) bool {
	maxY, maxX := len(es), len(es[0])
	for _, d := range dirs {
		x, y := c+d.dx, r+d.dy
		if x >= 0 && x < maxX && y >= 0 && y < maxY {
			v := es[y][x]
			if !unicode.IsDigit(v) && v != dot {
				return true
			}
		}
	}
	return false
}
