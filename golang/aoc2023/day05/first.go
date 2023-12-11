package day05

import (
	"fmt"
	"strings"
	"syodage/aoc2023/utils"
)

type SrcToDestMap struct {
	destStart int
	srcStart  int
	r         int
}

func (sd *SrcToDestMap) FindDest(src int) (int, bool) {
	if sd.srcStart <= src && src <= sd.srcStart+sd.r-1 {
		return sd.destStart + (src - sd.srcStart), true
	}
	return -1, false
}

func (sd *SrcToDestMap) Print() {
	fmt.Printf("[%d, %d, %d],", sd.destStart, sd.srcStart, sd.r)
}

type Almanac struct {
	Seeds        []int
	SeedToSoil   []*SrcToDestMap
	SoilToFert   []*SrcToDestMap
	FertToWater  []*SrcToDestMap
	WaterToLight []*SrcToDestMap
	LightToTemp  []*SrcToDestMap
	TempToHumid  []*SrcToDestMap
	HumidToLoc   []*SrcToDestMap
}

func (a *Almanac) FindLocations(seeds []int) ([]int, error) {
	// seed -> soil --> fertilizer -> water -> Light -> Temperature -> Humidity -> Location

	// seed -> soil
	soils := FindDestinations(seeds, a.SeedToSoil)
	// fmt.Printf("Soils: %+v\n", soils)
	fertilizers := FindDestinations(soils, a.SoilToFert)
	waters := FindDestinations(fertilizers, a.FertToWater)
	lights := FindDestinations(waters, a.WaterToLight)
	temps := FindDestinations(lights, a.LightToTemp)
	h := FindDestinations(temps, a.TempToHumid)
	locs := FindDestinations(h, a.HumidToLoc)
	return locs, nil
}

func FindDestinations(srcs []int, m []*SrcToDestMap) []int {
	var dests []int
	for _, s := range srcs {
		found := false
		for _, mm := range m {
			d, ok := mm.FindDest(s)
			if ok {
				dests = append(dests, d)
				found = true
				break
			}
		}
		// no dest find, use s as destination
		if !found {
			dests = append(dests, s)
		}
	}
	return dests
}

func (a *Almanac) Print() {
	// seed -> soil --> fertilizer -> water -> Light -> Temperature -> Humidity -> Location
	fmt.Printf("Seeds: %+v\n", a.Seeds)
	fmt.Println("SeedToSoil:")
	for _, ss := range a.SeedToSoil {
		ss.Print()
	}
	fmt.Println()
	fmt.Println("SoilToFert:")
	for _, ss := range a.SoilToFert {
		ss.Print()
	}
	fmt.Println()
	fmt.Println("FertToWater:")
	for _, ss := range a.FertToWater {
		ss.Print()
	}
	fmt.Println()
	fmt.Println("WaterToLight:")
	for _, ss := range a.WaterToLight {
		ss.Print()
	}
	fmt.Println()
	fmt.Println("LightToTemp:")
	for _, ss := range a.LightToTemp {
		ss.Print()
	}
	fmt.Println()
	fmt.Println("TempToHumid:")
	for _, ss := range a.TempToHumid {
		ss.Print()
	}
	fmt.Println()
	fmt.Println("HumidToLoc:")
	for _, ss := range a.HumidToLoc {
		ss.Print()
	}
	fmt.Println("\nEnd")
}

func FirstPart(inputs []string) (int, error) {
	alm, err := parseInput(inputs)
	if err != nil {
		return 0, err
	}
	// fmt.Printf("Alm: %#v\n", alm)
	// alm.Print()
	locs, err := alm.FindLocations(alm.Seeds)
	if err != nil {
		return 0, err
	}
	lowestLoc := locs[0]
	for _, loc := range locs[1:] {
		if lowestLoc > loc {
			lowestLoc = loc
		}
	}

	return lowestLoc, nil
}

func parseInput(inputs []string) (Almanac, error) {
	alm := Almanac{}

	i := 0
	for i < len(inputs) {
		input := inputs[i]
		if strings.HasPrefix(input, "seeds") {
			ss := strings.Split(input, " ")
			seed, err := utils.ToInts(ss[1:])
			if err != nil {
				return alm, err
			}
			alm.Seeds = append(alm.Seeds, seed...)
		}

		if strings.HasPrefix(input, "seed-to") {
			i += 1
			for i < len(inputs) {
				line := inputs[i]
				if line == "" {
					break
				}

				in, err := utils.ToInts(strings.Split(line, " "))
				if err != nil {
					return alm, err
				}
				alm.SeedToSoil = append(alm.SeedToSoil, &SrcToDestMap{in[0], in[1], in[2]})
				i += 1
			}
		}
		if strings.HasPrefix(input, "soil-to") {
			i += 1
			for i < len(inputs) {
				line := inputs[i]
				if line == "" {
					break
				}

				in, err := utils.ToInts(strings.Split(line, " "))
				if err != nil {
					return alm, err
				}
				alm.SoilToFert = append(alm.SoilToFert, &SrcToDestMap{in[0], in[1], in[2]})
				i += 1
			}
		}
		if strings.HasPrefix(input, "fert") {
			i += 1
			for i < len(inputs) {
				line := inputs[i]
				if line == "" {
					break
				}

				in, err := utils.ToInts(strings.Split(line, " "))
				if err != nil {
					return alm, err
				}
				alm.FertToWater = append(alm.FertToWater, &SrcToDestMap{in[0], in[1], in[2]})
				i += 1
			}
		}
		if strings.HasPrefix(input, "water") {
			i += 1
			for i < len(inputs) {
				line := inputs[i]
				if line == "" {
					break
				}

				in, err := utils.ToInts(strings.Split(line, " "))
				if err != nil {
					return alm, err
				}
				alm.WaterToLight = append(alm.WaterToLight, &SrcToDestMap{in[0], in[1], in[2]})
				i += 1
			}
		}
		if strings.HasPrefix(input, "light") {
			i += 1
			for i < len(inputs) {
				line := inputs[i]
				if line == "" {
					break
				}

				in, err := utils.ToInts(strings.Split(line, " "))
				if err != nil {
					return alm, err
				}
				alm.LightToTemp = append(alm.LightToTemp, &SrcToDestMap{in[0], in[1], in[2]})
				i += 1
			}
		}
		if strings.HasPrefix(input, "temp") {
			i += 1
			for i < len(inputs) {
				line := inputs[i]
				if line == "" {
					break
				}

				in, err := utils.ToInts(strings.Split(line, " "))
				if err != nil {
					return alm, err
				}
				alm.TempToHumid = append(alm.TempToHumid, &SrcToDestMap{in[0], in[1], in[2]})
				i += 1
			}
		}
		if strings.HasPrefix(input, "humidity") {
			i += 1
			for i < len(inputs) {
				line := inputs[i]
				if line == "" {
					break
				}

				in, err := utils.ToInts(strings.Split(line, " "))
				if err != nil {
					return alm, err
				}
				alm.HumidToLoc = append(alm.HumidToLoc, &SrcToDestMap{in[0], in[1], in[2]})
				i += 1
			}
		}

		i += 1
	}

	return alm, nil
}
