package day02

import "fmt"

func Second(inputs []string) (int, error) {
	// implement
	powers := 0
	for r, input := range inputs {
		game, err := parseInput(input)
		if err != nil {
			return 0, fmt.Errorf("Line: %d, error: %w", r, err)
		}

		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, set := range game.Sets {
			maxRed = max(maxRed, set.Red)
			maxBlue = max(maxBlue, set.Blue)
			maxGreen = max(maxGreen, set.Green)
		}

		powers += (maxRed * maxBlue * maxGreen)
	}
	return powers, nil
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
