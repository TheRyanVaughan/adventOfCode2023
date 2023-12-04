package main

import (
	"aoc2023/utils"
	"log"
)

const textFile = "day3/input.txt"

func main() {
	lines := utils.ReadTextFromFileToStringSlice(textFile)

	log.Println("Part 1: ", solvePart1(lines))
	// we have to reset lines btwn parts, since i'm overwriting it directly
	lines = utils.ReadTextFromFileToStringSlice(textFile)
	log.Println("Part 2: ", solvePart2(lines))
}

func solvePart1(lines []string) int64 {
	var sum int64
	for r := range lines {
		for c := range lines[r] {
			if isSymbol(lines[r][c]) {
				parts := getAdjacentParts(Position{r, c}, lines)

				for _, p := range parts {
					sum += p
				}
			}
		}
	}
	return sum
}

func solvePart2(lines []string) int64 {
	var sum int64
	for r := range lines {
		for c := range lines[r] {
			// just the symbol for a gear
			if isGear(lines[r][c]) {
				parts := getAdjacentParts(Position{r, c}, lines)
				// to confirm it is actually a gear, it should only have 2 parts
				if len(parts) == 2 {
					sum += parts[0] * parts[1]
				}
			}
		}
	}

	return sum
}

func getAdjacentParts(p Position, lines []string) []int64 {
	rows := len(lines)
	cols := len(lines[0])
	neighbors := getNeighbors(p)
	var parts []int64

	for _, n := range neighbors {
		if isInbounds(n, rows, cols) && isNumber(lines[n.row][n.col]) {
			parts = append(parts, parseNumber(lines, n.row, n.col))
		}
	}
	return parts
}

func getNeighbors(p Position) []Position {
	var neighbors [8]Position
	rowDeltas := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	colDeltas := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := range rowDeltas {
		neighbors[i] = Position{row: p.row + rowDeltas[i], col: p.col + colDeltas[i]}
	}
	return neighbors[:]
}

type Position struct {
	row int
	col int
}

// this should probably not modify the lines directly. perhaps a map to track that a position was already marked/used
func parseNumber(lines []string, row, col int) int64 {
	i := col
	for i > 0 && isNumber(lines[row][i-1]) {
		i--
	}

	var sum int64
	for i < len(lines) && isNumber(lines[row][i]) {
		sum *= 10
		sum += int64(lines[row][i] - '0')

		out := []rune(lines[row])
		out[i] = '.'
		lines[row] = string(out)

		i++
	}

	return sum
}

func isInbounds(p Position, rows int, cols int) bool {
	return p.row >= 0 && p.col >= 0 && p.row < rows && p.col < cols
}

func isSymbol(c byte) bool {
	return c != '.' && !isNumber(c)
}

func isGear(c byte) bool {
	return c == '*'
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}
