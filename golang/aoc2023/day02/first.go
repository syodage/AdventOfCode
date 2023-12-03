package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type Set struct {
	Green int
	Blue  int
	Red   int
}

type Game struct {
	Id   int
	Sets []*Set
}

func First(inputs []string) (int, error) {
	maxRed, maxGreen, maxBlue := 12, 13, 14
	sumGameIds := 0
	for r, input := range inputs {
		game, err := parseInput(input)
		if err != nil {
			return 0, fmt.Errorf("Line: %d, error: %w", r, err)
		}
		valid := true
		for _, set := range game.Sets {
			if set.Red > maxRed || set.Blue > maxBlue || set.Green > maxGreen {
				valid = false
				break
			}
		}
		if valid {
			sumGameIds += game.Id
		}
	}
	return sumGameIds, nil
}

func parseInput(input string) (*Game, error) {
	spaceSep := func(c rune) bool {
		return c == ' '
	}
	// fmt.Printf("Fields are: %q", strings.FieldsFunc(" 3 blue   ", f))
	sp := strings.Split(input, ":")
	if len(sp) != 2 {
		return nil, fmt.Errorf("Expected two parts but found %d", len(sp))
	}
	spGame := strings.FieldsFunc(sp[0], spaceSep)
	id, err := strconv.Atoi(spGame[1])
	if err != nil {
		return nil, fmt.Errorf("Error converting game id split %v, to int", spGame)
	}

	g := &Game{
		Id: id,
	}

	spSets := strings.Split(sp[1], ";")
	for _, spSet := range spSets {
		s := &Set{}
		for _, spCol := range strings.Split(spSet, ",") {
			spC := strings.FieldsFunc(spCol, spaceSep)
			if len(spC) != 2 {
				return nil, fmt.Errorf("Expected len 2 found len %d with array: %+v", len(spC), spC)
			}
			val, err := strconv.Atoi(spC[0])
			if err != nil {
				return nil, fmt.Errorf("Error converting color value: %s to int", spC[0])
			}
			switch spC[1] {
			case "green":
				s.Green = val
			case "red":
				s.Red = val
			case "blue":
				s.Blue = val
			default:
				return nil, fmt.Errorf("Invalid color: %s", spC[1])
			}
		}
		g.Sets = append(g.Sets, s)
	}
	return g, nil
}
