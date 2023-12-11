package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type MapRange struct {
	DestStart   int
	SourceStart int
	Range       int
}

type Map struct {
	Ranges []MapRange
}

type SeedToLocationMapper struct {
	Maps []Map
}

// Given int `n`, return true if `n` is within this maps source range
func (r MapRange) SourceInRange(n int) bool {
	minInclusive := r.SourceStart
	maxExclusive := r.SourceStart + r.Range
	if n >= minInclusive && n < maxExclusive {
		fmt.Printf("n: %v is within %d and %d\n", n, minInclusive, maxExclusive)
		return true
	}
	fmt.Printf("n: %v not within %d and %d\n", n, minInclusive, maxExclusive)
	return false
}

// Given a number `n`, map it to it's corresponding output value according
// to the list of `Ranges`
func (m Map) Map(n int) int {
	fmt.Printf("m: %+v\n", m)
	for _, r := range m.Ranges {
		fmt.Printf("r: %+v\n", r)
		if r.SourceInRange(n) {
			// convert
			destIndex := n - r.SourceStart
			return r.DestStart + destIndex
		}
	}

	// default to same number
	return n
}

// Given a list of seeds, maps each one all the way through to the corresponding
// location number, returning a list of locations for each seed.
func (sm SeedToLocationMapper) MapAll(seeds []int) (out []int) {
	for _, seed := range seeds {
		fmt.Printf("SEED: %d\n", seed)
		res := seed
		for _, m := range sm.Maps {
			fmt.Printf("res: %v\n", res)
			res = m.Map(res)
		}
		out = append(out, res)
	}
	return
}

const (
	NONE                    = -1
	SEEDS                   = 0
	SEED_TO_SOIL            = 1
	SOIL_TO_FERTILIZER      = 2
	FERTILIZER_TO_WATER     = 3
	WATER_TO_LIGHT          = 4
	LIGHT_TO_TEMPERATURE    = 5
	TEMPERATURE_TO_HUMIDITY = 6
	HUMIDITY_TO_LOCATION    = 7
)

func ParseInput(lines []string) (seeds []int, m SeedToLocationMapper) {
	mode := SEEDS
	lastMode := SEEDS
	var ranges []MapRange

	for _, line := range lines {
		// blank lines reset mode to -1
		if line == "" {
			// add all gathered ranges to a new map on the mapper
			if mode > SEEDS {
				m.Maps = append(m.Maps, Map{Ranges: ranges})
			}

			// clear ranges
			ranges = []MapRange{}

			// reset mode
			lastMode = mode
			mode = NONE
			continue
		}

		// check for next mode
		if mode == NONE {
			if strings.Contains(line, "map:") {
				mode = lastMode + 1
				continue
			}
		}

		// parse seeds
		if mode == SEEDS {
			seedStrs := strings.Split(strings.Replace(line, "seeds: ", "", -1), " ")
			for _, seedStr := range seedStrs {
				seedNo, err := strconv.Atoi(seedStr)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, seedNo)
			}
		}

		// parse maps
		if mode > SEEDS {
			ranges = append(ranges, ParseMapRange(line))
		}
	}

	// append final range since it wont be caught by blank line check
	m.Maps = append(m.Maps, Map{Ranges: ranges})

	return
}

func ParseMapRange(line string) MapRange {
	strs := strings.Split(line, " ")
	d, _ := strconv.Atoi(strs[0])
	s, _ := strconv.Atoi(strs[1])
	r, _ := strconv.Atoi(strs[2])
	return MapRange{
		DestStart:   d,
		SourceStart: s,
		Range:       r,
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	seeds, mapper := ParseInput(lines)

	fmt.Printf("seeds: %v\n", seeds)
	fmt.Printf("mapper: %+v\n", mapper)

	locations := mapper.MapAll(seeds)
	fmt.Printf("locations: %v\n", locations)

	minLocation := slices.Min(locations)
	fmt.Printf("minLocation: %v\n", minLocation)
}
