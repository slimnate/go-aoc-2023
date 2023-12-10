package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

/*
Steps Part 1:

Create a matrix of each character (2D array most likely)
	Create a list that has each possible symbol during parsing of data

Create a list of numbers that includes their start and end coordinates
	Store each number in a struct that includes it's value as well as start/end coords

Create a function that checks each number for any adjacent symbols
	Generate a list of all coordinates adjacent to the number
		same row: startCoord[row][col-1], endCoord[row][col+1]            - comma means two discreet values
		row above: startCoord[row-1][col-1] ... endCoord[row-1][col+1]    - ellipses means range of values
		row below: startcoord[row+1][col-1] ... endcoord[row+1][col+1]    - ellipses means range of values
		Be sure to check for any out of bounds coords and omit them from the generated list
	Check each coordinate in the list for a symbol
	If a symbol is found, add the number to a list of part numbers

Sum the list of part numbers


Part 2:

Store the location of each star
Loop through each star, checking the list of positioned numbers for a match, storing each match. If mathces == 2, it's a gear
*/

type RuneSlice []rune

func (s RuneSlice) Contains(c rune) bool {
	for _, r := range s {
		if c == r {
			return true
		}
	}
	return false
}

type PosSlice []Position

func (s PosSlice) Contains(p Position) bool {
	for _, i := range s {
		if p == i {
			return true
		}
	}
	return false
}

var nonSymbols = RuneSlice{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}

type Matrix struct {
	Height int
	Width  int
	Items  [][]rune
}

type Position struct {
	Row int
	Col int
}

type PositionedNumber struct {
	Value    int
	Start    Position
	End      Position
	adjacent []Position
}

func (m *Matrix) Set(row int, col int, val rune) {
	m.Items[row][col] = val
}

func (m *Matrix) Get(row int, col int) rune {
	return m.Items[row][col]
}

func (m *Matrix) Sprint() string {
	res := ""
	for row := 0; row < m.Height; row++ {
		for col := 0; col < m.Width; col++ {
			res += fmt.Sprint(string(m.Items[row][col]), " ")
		}
		res += "\n"
	}
	return res
}

func NewMatrix(width int, height int) *Matrix {
	items := make([][]rune, height)
	for i := 0; i < height; i++ {
		items[i] = make([]rune, width)
	}
	return &Matrix{
		Height: height,
		Width:  width,
		Items:  items,
	}
}

func GenerateMatrixFromLines(lines []string) *Matrix {
	height := len(lines)
	width := len(lines[0])

	matrix := NewMatrix(width, height)

	for row, line := range lines {
		for col, char := range line {
			matrix.Set(row, col, char)
		}
	}
	return matrix
}

func GenerateSymbols(lines []string) RuneSlice {
	var sym RuneSlice

	for _, line := range lines {
		for _, char := range line {
			if !nonSymbols.Contains(char) {
				sym = append(sym, char)
			}
		}
	}

	return sym
}

func ExtractNumbers(s string, row int) []PositionedNumber {
	var res []PositionedNumber
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllIndex([]byte(s), -1)
	for _, match := range matches {
		startIndex := match[0]
		endIndex := match[1]
		valueStr := s[startIndex:endIndex]
		value, _ := strconv.Atoi(valueStr)

		res = append(res, PositionedNumber{
			Start: Position{
				Row: row,
				Col: startIndex,
			},
			End: Position{
				Row: row,
				Col: endIndex - 1,
			},
			Value: value,
		})
	}
	return res
}

// Generate a list of all coordinates adjacent to the number
//
//	same row: startCoord[row][col-1], endCoord[row][col+1]            - comma means two discreet values
//	row above: startCoord[row-1][col-1] ... endCoord[row-1][col+1]    - ellipses means range of values
//	row below: startcoord[row+1][col-1] ... endcoord[row+1][col+1]    - ellipses means range of values
//	Be sure to check for any out of bounds coords and omit them from the generated list
func (p *PositionedNumber) AdjacentPositions(m Matrix) []Position {
	var res []Position

	if len(p.adjacent) == 0 {
		// same row before
		res = append(res, Position{Row: p.Start.Row, Col: p.Start.Col - 1})

		// same row after
		res = append(res, Position{Row: p.End.Row, Col: p.End.Col + 1})

		for col := p.Start.Col - 1; col <= p.End.Col+1; col++ {
			res = append(res, Position{Row: p.Start.Row - 1, Col: col})
			res = append(res, Position{Row: p.Start.Row + 1, Col: col})
		}

		p.adjacent = FilterValid(res, m.Width, m.Height)
	}

	return p.adjacent
}

func (p *PositionedNumber) IsAdjacent(pos Position, m *Matrix) bool {
	// fmt.Printf("p.adjacent: %v\n", p.adjacent)
	for _, adj := range p.AdjacentPositions(*m) {
		if adj.Col == pos.Col && adj.Row == pos.Row {
			// fmt.Printf("%+v is adjacent to %+v\n", pos, p)
			return true
		}
	}
	return false
}

func FilterValid(positions []Position, width int, height int) []Position {
	var valid []Position

	for _, p := range positions {
		if p.Col >= 0 && p.Col < width && p.Row >= 0 && p.Row < height {
			valid = append(valid, p)
		}
	}

	return valid
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	// read lines
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line: %v\n", line)
		lines = append(lines, line)
	}
	fmt.Println(lines)

	symbols := GenerateSymbols(lines)

	// generate matrix
	matrix := GenerateMatrixFromLines(lines)

	// generate list of numbers
	var nums []PositionedNumber
	for row := 0; row < matrix.Height; row++ {
		rowText := matrix.Items[row]
		// fmt.Print(string(rowText), "\n")
		nums = append(nums, ExtractNumbers(string(rowText), row)...)
	}

	fmt.Printf("matrix: %+v\n", matrix)

	// create parts list
	var partsList []PositionedNumber
	var starsList PosSlice
	for _, n := range nums {
		adjacent := n.AdjacentPositions(*matrix)
		// fmt.Printf("n: %+v\n", n)
		for _, a := range adjacent {
			char := matrix.Get(a.Row, a.Col)
			if symbols.Contains(char) {
				// add to parts list
				partsList = append(partsList, n)

				// check for a star and that we haven't already added it
				if char == '*' && !starsList.Contains(a) {
					starsList = append(starsList, a)
				}
			}
		}
	}
	// fmt.Printf("nums[0]: %+v\n", nums[0])

	// fmt.Printf("starsList: %v\n", starsList)

	var gearRatiosList []int
	for _, star := range starsList {
		var matches []PositionedNumber
		for _, num := range nums {
			// fmt.Printf("num: %+v  star: %+v\n", num, star)
			if num.IsAdjacent(star, matrix) {
				matches = append(matches, num)
			}
		}

		fmt.Printf("star: %+v\n", star)
		// fmt.Printf("matches: %+v\n", matches)
		fmt.Printf("len(matches): %v\n", len(matches))

		if len(matches) == 2 {
			fmt.Println(" adding gears ")
			// Exactly two matches means it's a gear
			gearRatiosList = append(gearRatiosList, matches[0].Value*matches[1].Value)
		}
	}

	fmt.Printf("gearRatiosList: %v\n", gearRatiosList)

	// sum gear ratios
	gearRatioSum := 0
	for _, ratio := range gearRatiosList {
		gearRatioSum += ratio
	}

	// sum parts list
	partsListSum := 0
	for _, part := range partsList {
		partsListSum += part.Value
	}

	fmt.Printf("partsListSum: %v\n", partsListSum)
	fmt.Printf("gearRatioSum: %v\n", gearRatioSum)
}
