package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Tracks the number of each color cube pulled in a single round of a game
type Round struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Id     int
	Rounds []*Round
}

func ParseGame(s string) (g *Game) {
	g = &Game{}

	split1 := strings.Split(s, ": ")

	// extract id from first part of string
	id, _ := strconv.Atoi(strings.Replace(split1[0], "Game ", "", 1))
	g.Id = id

	// Extract rounds from last part of string
	roundStrings := strings.Split(split1[1], "; ")
	for _, roundString := range roundStrings {
		// fmt.Printf("roundString: %v\n", roundString)
		g.Rounds = append(g.Rounds, ParseRound(roundString))
	}

	return
}

func ParseRound(s string) (r *Round) {
	r = &Round{}
	splits := strings.Split(s, ", ")

	for _, roundString := range splits {
		roundSplit := strings.Split(roundString, " ")
		count, _ := strconv.Atoi(roundSplit[0])
		color := roundSplit[1]
		if color == "red" {
			r.Red = count
		} else if color == "green" {
			r.Green = count
		} else if color == "blue" {
			r.Blue = count
		}
	}

	return
}

func (g *Game) IsPossible(red int, green int, blue int) bool {
	for _, round := range g.Rounds {
		if round.Red > red || round.Green > green || round.Blue > blue {
			return false
		}
	}
	return true
}

func (g *Game) MaxEachColor() (red int, green int, blue int) {
	for _, round := range g.Rounds {
		if round.Red > red {
			red = round.Red
		}
		if round.Green > green {
			green = round.Green
		}
		if round.Blue > blue {
			blue = round.Blue
		}
	}

	return
}

func main() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)

	var games []*Game
	for scanner.Scan() {
		line := scanner.Text()

		g := ParseGame(line)

		games = append(games, g)
	}

	idTotal := 0
	for _, game := range games {
		if game.IsPossible(12, 13, 14) {
			fmt.Printf("Game %d possible \n", game.Id)
			idTotal += game.Id
		}
	}

	powerTotal := 0
	for _, game := range games {
		r, g, b := game.MaxEachColor()
		power := r * g * b
		powerTotal += power
	}

	fmt.Printf("idTotal: %v\n", idTotal)
	fmt.Printf("powerTotal: %v\n", powerTotal)
}
