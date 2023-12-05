package day03

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Cell struct {
	y int
	x int
}

func Second(inputs []string) (int, error) {
	es := parseInput(inputs)
	num := strings.Builder{}
	m := make(map[Cell][]string)
	for r, line := range es {
		adjSym := map[Cell]bool{}
		for c, sym := range line {
			if !unicode.IsDigit(sym) {
				if num.Len() != 0 && len(adjSym) > 0 {
					rNum := num.String()
					for ad := range adjSym {
						_, ok := m[ad]
						if !ok {
							m[ad] = []string{}
						}
						m[ad] = append(m[ad], rNum)
					}
				}
				num.Reset()
				adjSym = make(map[Cell]bool)
				continue
			}

			if unicode.IsDigit(sym) {
				for _, sym := range adjacentSymbols(es, c, r) {
					_, ok := adjSym[sym]
					if !ok {
						adjSym[sym] = true
					}
				}
				num.WriteRune(sym)
			}
		}
		// if a valid number with adjacent to a Symbol
		if num.Len() != 0 && len(adjSym) > 0 {
			for ad := range adjSym {
				_, ok := m[ad]
				if !ok {
					m[ad] = []string{}
				}
				m[ad] = append(m[ad], num.String())
			}
		}
	}

	// fmt.Printf("Map: %+v\n", m)
	sumOfGearRatio := 0
	for k, nums := range m {
		i, j := k.y, k.x
		if rune(es[i][j]) == rune('*') && len(nums) == 2 {
			gearRatio := 1
			for _, nStr := range nums {
				n, err := strconv.Atoi(nStr)
				if err != nil {
					return 0, fmt.Errorf("Error converting %q a int, error: %v", nStr, err)
				}
				gearRatio *= n
			}
			sumOfGearRatio += gearRatio
		}
	}
	return sumOfGearRatio, nil
}

func adjacentSymbols(es [][]rune, c, r int) []Cell {
	maxY, maxX := len(es), len(es[0])
	var res []Cell
	for _, d := range dirs {
		x, y := c+d.dx, r+d.dy
		if x >= 0 && x < maxX && y >= 0 && y < maxY {
			v := es[y][x]
			if !unicode.IsDigit(v) && v != dot {
				res = append(res, Cell{y, x})
			}
		}
	}
	return res
}

// func findAstrixs(es [][]rune) []Cell {
// 	var res []Cell = nil
// 	for i, row := range es {
// 		for j, col := range row {
// 			if col == rune('*') {
// 				res = append(res, Cell{i, j})
// 			}
// 		}
// 	}
// 	return res
