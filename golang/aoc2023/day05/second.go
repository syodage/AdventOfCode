package day05

import (
	"fmt"
	"math"
	"sort"
)

func SecondPart(inputs []string) (int, error) {

	alm, err := parseInput(inputs)
	if err != nil {
		return 0, err
	}
	// fmt.Printf("Alm: %#v\n", alm)
	// alm.Print()
	seedRanges := alm.Seeds
	if len(seedRanges)%2 != 0 {
		return 0, fmt.Errorf("Invalid seed range input, len: %d, not pairable", len(alm.Seeds))
	}
	lowestLoc := math.MaxInt32
	// V1 start
	// iterate range by range
	// ranges := len(seedRanges) / 2
	// for r := 0; r < ranges; r++ {
	// 	// fmt.Printf("Seed range: [%d, %d]\n", seedRanges[r*2], seedRanges[r*2+1])
	// 	start := seedRanges[r*2]
	// 	end := start + seedRanges[r*2+1]
	// 	for j := seedRanges[r*2]; j < end; j++ {
	// 		ss := []int{j}
	// 		locs, err := alm.FindLocations(ss)
	// 		if err != nil {
	// 			return 0, err
	// 		}
	// 		if lowestLoc > locs[0] {
	// 			lowestLoc = locs[0]
	// 		}
	// 	}
	// }
	// V1 End

	// V2 Start
	uRanges := findUniqueSeedRanges(alm.Seeds)
	for _, r := range uRanges {
		for i := r.start; i <= r.end; i++ {
			ss := []int{i}
			locs, err := alm.FindLocations(ss)
			if err != nil {
				return 0, err
			}
			if lowestLoc > locs[0] {
				lowestLoc = locs[0]
			}

		}
	}
	// V2 End
	return lowestLoc, nil
}

type Range struct {
	start int
	end   int
}

func findUniqueSeedRanges(seeds []int) []*Range {
	var ranges []*Range
	rLen := len(seeds) / 2

	for r := 0; r < rLen; r++ {
		// fmt.Printf("Seed range: [%d, %d]\n", seedRanges[r*2], seedRanges[r*2+1])
		start := seeds[r*2]
		end := start + seeds[r*2+1] - 1
		ranges = append(ranges, &Range{start, end})
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start <= ranges[j].start
	})

	return mergeRanges(ranges)
}

func mergeRanges(r []*Range) []*Range {
	var mR []*Range

	if len(r) == 1 {
		return r[:]
	}

	cr := r[0]
	j := 1
	for j < len(r) {
		if cr.end < r[j].start {
			mR = append(mR, cr)
			cr = r[j]
		} else {
			cr.start = min(cr.start, r[j].start)
			cr.end = max(cr.end, r[j].end)
		}

		j++
	}
	// add cr
	mR = append(mR, cr)
	return mR
}
