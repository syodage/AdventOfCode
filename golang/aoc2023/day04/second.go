package day04

func SecondPart(inputs []string) (int, error) {
	cards := parseInputs(inputs)
	// initialized buckets
	buckets := make([]int, len(cards))
	for i := range cards {
		buckets[i] = 1
	}

	for i, c := range cards {
		numOfCards := buckets[i]
		for j := range c.MatchNums() {
			buckets[i+j+1] += numOfCards
		}
	}

	totalScratchCards := 0
	for _, nc := range buckets {
		totalScratchCards += nc
	}

	return totalScratchCards, nil
}
