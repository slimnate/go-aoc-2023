package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"time"
)

/*
Part 1:
The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

Steps:
Parse each line
Get the first and last numerical digits from each line
Concatenate to get a list of 2-digit numbers
Calculate the sum of these numbers

*/

var digits = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

var numMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"0":     0,
}

func main() {
	// open file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// scan lines
	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var first rune
		var last rune
		line := scanner.Text()

		// loop over chars, grabbing first and last numerical value
		for _, runeValue := range line {
			if slices.Contains(digits, runeValue) {
				if first == 0 {
					first = runeValue
				} else {
					last = runeValue
				}
			}
		}
		// if only one digit, first and last are same
		if last == 0 {
			last = first
		}

		// convert to number and add to numbers list
		num, err := strconv.Atoi(fmt.Sprint(string(first), string(last)))
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
	}

	// return sum of numbers
	var sum int
	for _, n := range numbers {
		sum += n
	}

	fmt.Printf("sum part 1: %v\n", sum)

	// PART 2
	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)

	// fmt.Printf("file: %+v\n", file)
	// fmt.Printf("scanner: %+v\n", scanner)

	scanCount := 0
	iters := 0
	startTime := time.Now()
	partTwoTotal := 0

	for scanner.Scan() {
		line := scanner.Text()

		runes := []rune(line)
		// fmt.Printf("line: %v\n", line)

		var nums []int

		// loop through each character of the line
		for i := 0; i < len(line); i++ {
			// fmt.Printf("i: %v\n", i)
			// check each possible substring starting at i for a number match
			for j := i + 1; j < len(line)+1; j++ {
				// fmt.Printf("j: %v\n", j)
				substr := string(runes[i:j])
				num, ok := numMap[substr]
				if ok {
					nums = append(nums, num)
				}
				iters++
				// fmt.Printf("substr: %v\n", substr)
			}
		}
		numStr := fmt.Sprint(strconv.Itoa(nums[0]), strconv.Itoa(nums[len(nums)-1]))
		n, _ := strconv.Atoi(numStr)
		partTwoTotal += n

		// fmt.Printf("iters: %v\n", iters)
		// fmt.Printf("nums: %v\n", nums)

		scanCount++
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	timePerLine := float64(duration.Milliseconds()) / float64(scanCount)
	fmt.Printf("partTwoTotal: %v\n", partTwoTotal)
	fmt.Printf("duration: %v\n", duration)
	fmt.Printf("timePerLine: %v ms\n", timePerLine)
}

/*

With print substr:
duration: 16.121139s
timePerLine: 16 ms


Without print substr:
duration: 192.7726ms
timePerLine: 0.192 ms

duration: 192.2135ms
timePerLine: 0.192 ms

duration: 195.8547ms
timePerLine: 0.195 ms

duration: 194.4238ms
timePerLine: 0.194 ms


Without any prints:

*/
