package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
		// fmt.Printf("n: %v is within %d and %d\n", n, minInclusive, maxExclusive)
		return true
	}
	// fmt.Printf("n: %v not within %d and %d\n", n, minInclusive, maxExclusive)
	return false
}

// Given a number `n`, map it to it's corresponding output value according
// to the list of `Ranges`
func (m Map) Map(n int) int {
	// fmt.Printf("m: %+v\n", m)
	for _, r := range m.Ranges {
		// fmt.Printf("r: %+v\n", r)
		if r.SourceInRange(n) {
			// convert
			val := r.DestStart + (n - r.SourceStart)
			// fmt.Printf("%v > %v\n", n, val)
			return val
		}
	}

	// default to same number
	// fmt.Printf("%v > %v\n", n, n)
	return n
}

// Given a list of seeds, maps each one all the way through to the corresponding
// location number, returning a list of locations for each seed.
func (sm SeedToLocationMapper) MapAll(seeds []int) (out []int) {
	for _, seed := range seeds {
		fmt.Printf("SEED: %d\n", seed)
		res := seed
		for _, m := range sm.Maps {
			res = m.Map(res)
			// fmt.Printf(" > %v", res)
		}
		// fmt.Println("")
		out = append(out, res)
	}
	return
}

// Given a list of seeds, maps each one all the way to the corresponding
// location number, returning the minimum location value
func (sm SeedToLocationMapper) MapMinPairs(seeds []int) (min int) {
	var start int
	min = math.MaxInt64 // min starts at the maximum possible int value

	for i, n := range seeds {
		if i == 0 || i%2 == 0 {
			// first pair is the start number
			start = n
		} else {
			// second number is the range length
			for v := start; v < start+n; v++ {
				res := v
				// for each number in the range, map them down to the location number
				// fmt.Printf("seed: %v", v)
				for _, m := range sm.Maps {
					res = m.Map(res)
					// fmt.Printf(" > %v", res)
				}
				// fmt.Println("")
				// check for a new min
				if res < min {
					min = res
				}
			}
		}
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

const (
	SEED_MODE_NORMAL = 1
	SEED_MODE_PAIRS  = 2
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
			seeds = ParseSeeds(line)
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

func ParseSeeds(line string) (seeds []int) {
	seedStrs := strings.Split(strings.Replace(line, "seeds: ", "", -1), " ")
	for _, seedStr := range seedStrs {
		seedNo, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seedNo)
	}
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

	fmt.Printf("seeds: %v\n", len(seeds))
	fmt.Printf("mapper: %+v\n", mapper)

	// locations := mapper.MapAll(seeds)
	// fmt.Printf("locations: %v\n", locations)

	minLocation := mapper.MapMinPairs(seeds)
	fmt.Printf("minLocation: %v\n", minLocation)
}
