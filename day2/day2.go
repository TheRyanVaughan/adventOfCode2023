package main

import (
	"aoc2023/utils"
	"log"
	"strconv"
	"strings"
)

const inputFile = "day2/sample.txt"

func main() {
	lines := utils.ReadTextFromFileToStringSlice(inputFile)

	log.Println("Part 1: ", solvePart1(lines))
	log.Println("Part 2: ", solvePart2(lines))
}

const (
	maxRed = iota + 12
	maxGreen
	maxBlue
)

type game struct {
	id      int64
	matches []match
}

type match struct {
	red   int64
	blue  int64
	green int64
}

func parseMatch(draw string) match {
	var m match
	for _, countColor := range strings.Split(draw, ",") {
		countColor = strings.TrimSpace(countColor)
		cc := strings.Split(countColor, " ")
		count, _ := strconv.ParseInt(cc[0], 10, 64)
		if cc[1][0] == 'r' {
			m.red = count
		} else if cc[1][0] == 'g' {
			m.green = count
		} else {
			m.blue = count
		}
	}
	return m
}
func parseIntoGames(lines []string) []game {
	games := make([]game, len(lines))
	for i, line := range lines {
		var currGame game
		split := strings.SplitN(line, ":", 2)
		gameId := split[0]
		// gameId is of format "Game id"
		id, _ := strconv.ParseInt(strings.Split(gameId, " ")[1], 10, 64)

		currGame.id = id

		matches := split[1]
		for _, draw := range strings.Split(matches, ";") {
			m := parseMatch(draw)

			currGame.matches = append(currGame.matches, m)
		}
		games[i] = currGame
	}
	return games
}
func solvePart1(lines []string) int {
	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	var sum int
	games := parseIntoGames(lines)

	for _, game := range games {
		possible := true
		for _, match := range game.matches {
			if match.red > maxRed || match.blue > maxBlue || match.green > maxGreen {
				possible = false
			}
		}

		if possible {
			sum += int(game.id)
		}
	}

	return sum
}

func solvePart2(lines []string) int {
	var sum int
	games := parseIntoGames(lines)

	for _, game := range games {
		var max match
		for _, m := range game.matches {
			max.red = maxInt(max.red, m.red)
			max.green = maxInt(max.green, m.green)
			max.blue = maxInt(max.blue, m.blue)
		}
		sum += int(max.red) * int(max.green) * int(max.blue)
	}

	return sum
}

func maxInt(a, b int64) int64 {
	if a > b {
		return int64(a)
	}
	return int64(b)
}
