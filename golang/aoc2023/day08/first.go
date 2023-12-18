package day08

import (
	"fmt"
	"syodage/aoc2023/utils"
)

type Network struct {
	LeftNode  string
	RightNode string
}

func (n *Network) String() string {
	return fmt.Sprintf("{left: %q, right: %q}", n.LeftNode, n.RightNode)
}

func FirstPart(inputs []string) (int, error) {
	upperBound := 100000000
	dirs, network := parseInputs(inputs)
	// fmt.Println(dirs)
	// fmt.Printf("map: %+v\n", network)
	step := 0
	ln := len(dirs)
	currNode := "AAA"
	for {
		r := rune(dirs[step%ln])
		step += 1
		node, ok := network[currNode]
		if !ok {
			return 0, fmt.Errorf("currNode: %q, not found", currNode)
		}
		if r == rune('L') {
			currNode = node.LeftNode
		}

		if r == rune('R') {
			currNode = node.RightNode
		}
		// fmt.Printf("CurrNode: %q\n", currNode)
		if currNode == "ZZZ" {
			return step, nil
		}
		if step >= upperBound {
			return step, fmt.Errorf("step exceeds upper bound: %d", upperBound)
		}
	}
}

func parseInputs(inputs []string) (string, map[string]*Network) {
	dir := inputs[0]

	nMap := make(map[string]*Network)
	for _, in := range inputs[2:] {
		sp := utils.SplitBySpaces(in)
		// fmt.Printf("sp: %+v\n", strings.Join(sp, "#"))
		nMap[sp[0]] = &Network{sp[2][1:4], sp[3][:3]}
	}
	return dir, nMap
}
