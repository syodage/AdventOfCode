package day11

import "fmt"

const (
	Galaxy = rune('#')
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("{C: %d, R: %d}", p.X, p.Y)
}

func FirstPart(inputs []string, spaceExpandFactor int) (int, error) {

	galaxyRows := make(map[int]bool)
	galaxyCols := make(map[int]bool)
	galaxies := []Point{}

	for r, row := range inputs {
		foundGalaxy := false
		for c, col := range row {
			if col == Galaxy {
				foundGalaxy = true
				galaxies = append(galaxies, Point{c, r})
			}
		}
		if foundGalaxy {
			galaxyRows[r] = true
		}
	}

	for c := 0; c < len(inputs[0]); c++ {
		foundGalaxy := false
		for r := 0; r < len(inputs); r++ {
			if Galaxy == rune(inputs[r][c]) {
				foundGalaxy = true
				break
			}
		}
		if foundGalaxy {
			galaxyCols[c] = true
		}
	}

	// fmt.Printf("Galaxy Rows: %+v\n", galaxyRows)
	// fmt.Printf("Galaxy Cols: %+v\n", galaxyCols)
	// fmt.Printf("Galaxies: %+v\n", galaxies)

	totalDist := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist := 0
			f, s := galaxies[i], galaxies[j]
			// Horizontal distance
			start := min(f.X, s.X)
			end := max(f.X, s.X)
			for x := start; x < end; x++ {
				if _, ok := galaxyCols[x]; !ok {
					dist += 1 * spaceExpandFactor
					continue
				}
				dist += 1
			}
			// Vertical distance
			start = min(f.Y, s.Y)
			end = max(f.Y, s.Y)
			for y := start; y < end; y++ {
				if _, ok := galaxyRows[y]; !ok {
					dist += 1 * spaceExpandFactor
					continue
				}

				dist += 1
			}

			// fmt.Printf("%v -> %v: %d", f, s, dist)
			totalDist += dist
		}
	}
	return totalDist, nil
}
