package day08

import (
	"fmt"
	"strings"
)

func SecondPart(inputs []string) (int, error) {
	upperBound := 100000000000
	dirs, network := parseInputs(inputs)
	// fmt.Println(dirs)
	// fmt.Printf("map: %+v\n", network)
	currNodes := make(map[int]string)
	count := 0
	for k := range network {
		if strings.HasSuffix(k, "A") {
			currNodes[count] = k
			count += 1
		}
	}
	fmt.Printf("Starting nodes count: %d, nodes: %+v\n", len(currNodes), currNodes)
	step := 0
	ln := len(dirs)
	count = 0
	for {
		r := rune(dirs[step%ln])
		step += 1
		nextNodes := make(map[int]string)
		for i, curr := range currNodes {
			node, ok := network[curr]
			if !ok {
				return 0, fmt.Errorf("curr: %q, not found", curr)
			}
			if r == rune('L') {
				nextNodes[i] = node.LeftNode
			}

			if r == rune('R') {
				nextNodes[i] = node.RightNode
			}
		}
		// if step < 100 {
		// 	fmt.Printf("Next nodes: %+v\n", nextNodes)
		//
		// } else {
		// 	return step, fmt.Errorf("Temp exit!")
		// }

		endsWithZ := 0
		for _, v := range nextNodes {
			if strings.HasSuffix(v, "Z") {
				endsWithZ += 1
			}
		}
		if endsWithZ > 2 {
			fmt.Printf("Count: %d, Step: %d, # of Z: %d, nodes: %+v\n", count, step, endsWithZ, nextNodes)
			count += 1
		}
		if endsWithZ == len(nextNodes) {
			return step, nil
		}
		currNodes = nextNodes

		if step >= upperBound {
			return step, fmt.Errorf("step exceeds upper bound: %d", upperBound)
		}
	}
}
