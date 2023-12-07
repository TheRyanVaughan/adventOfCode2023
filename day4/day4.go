package main

import (
	"aoc2023/utils"
	"strconv"
	"strings"
)

func main() {
	utils.SolveWith(4, solvePart1, solvePart2, parse, true)
}

type numMap map[int]bool

func solvePart1(cards []card) int {
	var sum int
	for _, c := range cards {
		u := len(c.numbers.intersect(c.winners))
		if u > 0 {
			sum += 1 << (u - 1)
		}
	}
	return sum
}

func solvePart2(cards []card) int {
	var cardCounts = make(map[int]int)
	var sum int

	// start with 1 of each card
	for i := range cards {
		cardCounts[i] = 1
	}

	for i, c := range cards {
		u := len(c.numbers.intersect(c.winners))
		for v := 1; v <= u; v++ {
			cardCounts[i+v] += cardCounts[i]
		}
		sum += cardCounts[i]
	}
	return sum
}
func (a numMap) intersect(b numMap) numMap {
	s := make(numMap)
	for k := range a {
		if b[k] {
			s[k] = true
		}
	}
	return s
}

type card struct {
	id      int
	winners numMap
	numbers numMap
}

func parse(lines []string) []card {
	var cards = make([]card, len(lines))
	for i, v := range lines {
		id := strings.Split(strings.TrimSpace(strings.Split(v, ":")[0]), " ")[1]
		intId, _ := strconv.ParseInt(id, 10, 32)
		info := strings.TrimSpace(strings.Split(v, ":")[1])
		// only two digit numbers in this problem
		info = strings.ReplaceAll(info, "  ", " ")
		numLists := strings.Split(info, "|")

		winners, numbers := strings.TrimSpace(numLists[0]), strings.TrimSpace(numLists[1])

		cards[i].id = int(intId)
		cards[i].winners = parseToCard(winners)
		cards[i].numbers = parseToCard(numbers)
	}
	return cards
}

func parseToCard(nums string) numMap {
	ints := make(numMap)
	for _, v := range strings.Split(nums, " ") {
		n, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			panic(err)
		}
		ints[int(n)] = true
	}
	return ints
}
