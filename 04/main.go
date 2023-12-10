package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntSlice []int

type Card struct {
	CardNumber     int
	WinningNumbers IntSlice
	MyNumbers      IntSlice
}

func (s IntSlice) Contains(v int) bool {
	for _, n := range s {
		if n == v {
			return true
		}
	}
	return false
}

func ParseCard(line string) Card {
	s1 := strings.Split(line, ": ")
	s2 := strings.Split(s1[0], " ")
	cardNo, _ := strconv.Atoi(s2[1])

	s3 := strings.Split(s1[1], " | ")
	s4 := strings.Split(s3[0], " ")
	s5 := strings.Split(s3[1], " ")

	var winningNos []int
	var myNos []int

	for _, winningNoStr := range s4 {
		if winningNoStr == "" {
			continue
		}
		winningNo, _ := strconv.Atoi(winningNoStr)
		winningNos = append(winningNos, winningNo)
	}

	for _, myNoStr := range s5 {
		if myNoStr == "" {
			continue
		}
		myNo, _ := strconv.Atoi(myNoStr)
		myNos = append(myNos, myNo)
	}

	return Card{
		CardNumber:     cardNo,
		WinningNumbers: winningNos,
		MyNumbers:      myNos,
	}
}

func (c Card) WinningNumberCount() int {
	winningNumberCount := 0
	for _, myNo := range c.MyNumbers {
		if c.WinningNumbers.Contains(myNo) {
			winningNumberCount++
		}
	}
	return winningNumberCount
}

func (c Card) Value() int {
	wnc := c.WinningNumberCount()
	if wnc == 0 {
		return 0
	}
	if wnc == 1 {
		return 1
	}
	return int(math.Pow(2, math.Max(0, float64(wnc-1))))
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var cards []Card
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, ParseCard(line))
	}

	totalValue := 0
	for _, card := range cards {
		totalValue += card.Value()
	}

	fmt.Printf("totalValue: %v\n", totalValue)

}
