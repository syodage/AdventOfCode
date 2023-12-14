package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func SplitBySpaces(in string) []string {
	return strings.FieldsFunc(in, RuneSep)
}

func ToInts(s []string) ([]int, error) {
	var nums []int
	for _, val := range s {
		n, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("Error converting %q to int, err: %v", val, err)
		}
		nums = append(nums, n)
	}
	return nums, nil
}
