package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func WaysToWin(raceTime int, record int) int {
	minWinTime := MinButtonTimeToWin(raceTime, record)
	maxWinTime := MaxButtonTimeToWin(raceTime, record)

	winningOptionCount := maxWinTime - minWinTime + 1

	fmt.Printf("minWinTime: %v\n", minWinTime)
	fmt.Printf("maxWinTime: %v\n", maxWinTime)
	fmt.Printf("winningOptionCount: %v\n", winningOptionCount)
	return winningOptionCount
}

func main() {
	lines, _ := utils.Lines("input.txt")

	numsRegex := regexp.MustCompile(`\d+`)

	timeStrs := numsRegex.FindAllString(lines[0], -1)
	distanceStrs := numsRegex.FindAllString(lines[1], -1)

	times, _ := utils.StrToIntSlice(timeStrs)
	records, _ := utils.StrToIntSlice(distanceStrs)

	winningOptionsProduct := 1

	// Part 1 - for each race, calculate the min and max button times that will beat the record
	for i, time := range times {
		winningOptionCount := WaysToWin(time, records[i])
		winningOptionsProduct *= winningOptionCount
	}

	fmt.Printf("winningOptionsProduct: %v\n", winningOptionsProduct)

	// Part 2 - same problem, but parsed differently, so there is only one race
	timesStr := strings.Split(lines[0], ":")[1]
	timesStr = strings.ReplaceAll(timesStr, " ", "")
	time, _ := strconv.Atoi(timesStr)

	recrodsStr := strings.Split(lines[1], ":")[1]
	recrodsStr = strings.ReplaceAll(recrodsStr, " ", "")
	record, _ := strconv.Atoi(recrodsStr)

	winningOptionsCount := WaysToWin(time, record)

	fmt.Printf("winningOptionsCount: %v\n", winningOptionsCount)
}
