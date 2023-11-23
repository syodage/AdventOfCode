package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day01() error {
	filePath := "./day01Input.txt"
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error reading input file: %s", filePath)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	count := 0 // Depth increases count
	// Assumption:  no depth with zero measurement.
	prevDepth := 0
	for scanner.Scan() {
		depthStr := scanner.Text()
		nextDepth, err := strconv.Atoi(depthStr)
		if err != nil {
			return fmt.Errorf("Error converting %s to int", depthStr)
		}
		// First measurement is not an increment, if prevDepth is zero don't increment the count.
		if prevDepth != 0 && prevDepth < nextDepth {
			count += 1
		}
		prevDepth = nextDepth
	}
	fmt.Printf("Depth Increases %d times.\n", count)
	return nil
}
