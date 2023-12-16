package day07

func SecondPart(inputs []string) (int, error) {

	cList, err := parseInput(inputs, true)
	if err != nil {
		return -1, err
	}

	SortCard(cList, true)
	// fmt.Printf("Cards:\n %+v\n", cList)
	totalWining := 0
	for i, c := range cList {
		// fmt.Printf("bid * rank = %d * %d", c.Bid, i+1)
		totalWining += (c.Bid * (i + 1))
	}
	return totalWining, nil
}
