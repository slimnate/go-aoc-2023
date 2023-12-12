package main

import (
	"fmt"
	"regexp"
	"utils"
)

func Distance(btnTime int, raceTime int) int {
	movingTime := raceTime - btnTime
	speed := btnTime
	distance := speed * movingTime
	return distance
}

func MinButtonTimeToWin(raceTime int, record int) int {
	for btnTime := 0; btnTime < raceTime; btnTime++ {
		if Distance(btnTime, raceTime) > record {
			return btnTime
		}
	}

	return -1
}

func MaxButtonTimeToWin(raceTime int, record int) int {
	for btnTime := raceTime; btnTime >= 0; btnTime-- {
		if Distance(btnTime, raceTime) > record {
			return btnTime
		}
	}

	return -1
}

func main() {
	lines, _ := utils.Lines("input.txt")

	numsRegex := regexp.MustCompile(`\d+`)

	timeStrs := numsRegex.FindAllString(lines[0], -1)
	distanceStrs := numsRegex.FindAllString(lines[1], -1)

	times, _ := utils.StrToIntSlice(timeStrs)
	records, _ := utils.StrToIntSlice(distanceStrs)

	winningOptionsProduct := 1

	// for each race, calculate the min and max button times that will beat the record
	for i, time := range times {
		minWinTime := MinButtonTimeToWin(time, records[i])
		maxWinTime := MaxButtonTimeToWin(time, records[i])

		winningOptionCount := maxWinTime - minWinTime + 1

		fmt.Printf("minWinTime: %v\n", minWinTime)
		fmt.Printf("maxWinTime: %v\n", maxWinTime)
		fmt.Printf("winningOptionCount: %v\n", winningOptionCount)

		winningOptionsProduct *= winningOptionCount
	}

	fmt.Printf("winningOptionsProduct: %v\n", winningOptionsProduct)
}
