package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading input file: %s", filePath)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	inputs := []string{}
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	return inputs, nil
}

func RuneSep(c rune) bool {
	return c == ' '
}
