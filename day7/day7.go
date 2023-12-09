package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

const runInput = true
const day = 7

type rank int

const (
	joker rank = iota - 1
	one
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

type level int

const (
	high level = iota
	onePair
	twoPair
	threeKind
	fullHouse
	fourKind
	fiveKind
)

func main() {
	var filePath string
	if runInput {
		filePath = fmt.Sprintf("day%d/input.txt", day)
	} else {
		filePath = fmt.Sprintf("day%d/sample.txt", day)
	}

	lines := utils.ReadTextFromFileToStringSlice(filePath)
	d := parse(lines)

	log.Println("Part 1: ", solvePart1(d))
	log.Println("Part 2: ", solvePart2(d))
}

func solvePart1(hands []hand) int {
	ret := 0

	order := hands[:]

	// reverse order
	sort.Slice(order, func(i, j int) bool {
		return order[i].compareTo(&order[j]) < 0
	})

	for i, v := range order {
		ret += (i + 1) * v.bid
	}
	return ret
}

func solvePart2(hands []hand) int {
	// i did not architecture this well since i put a lot of logic inside of my parser.
	// trick will be before sorting, recompute the hand value for everything
	// given new rules
	// thankfully, the best hand is always to replace jack with whichever card has greatest count

	order := hands[:]

	// transform the data
	for i, h := range order {
		keyCounts := make(map[rank]int)
		for j, v := range h.cards {
			if v == jack {
				order[i].cards[j] = joker
				v = joker
			}
			keyCounts[v] = keyCounts[v] + 1
		}

		var maxKey rank = -100 // invalid rank
		for k, v := range keyCounts {
			if k != joker && v > keyCounts[rank(maxKey)] {
				maxKey = k
			}
		}

		var newArr [5]rank = order[i].cards

		for j, v := range newArr {
			if v == joker {
				newArr[j] = maxKey
			}
		}

		newHandValue := getHandValue(newArr)
		order[i].handLevel = newHandValue
	}

	ret := 0
	// reverse order
	sort.Slice(order, func(i, j int) bool {
		return order[i].compareTo(&order[j]) < 0
	})

	for i, v := range order {
		log.Println(v)
		ret += (i + 1) * v.bid
	}
	return ret
}

type hand struct {
	handLevel level
	bid       int
	cards     [5]rank
}

func (first *hand) compareTo(second *hand) int {
	// returns -1 if first < second, 0 if equal, 1 if first > second
	if first.handLevel < second.handLevel {
		return -1
	} else if first.handLevel > second.handLevel {
		return 1
	} else {
		for i := range first.cards {
			if first.cards[i] < second.cards[i] {
				return -1
			} else if first.cards[i] > second.cards[i] {
				return 1
			}
		}
		log.Println("This shouldn't have happened", first, second)
		return 0
	}
}

func parse(lines []string) []hand {
	hands := make([]hand, len(lines))

	for i, line := range lines {
		s := strings.Split(line, " ")

		bid, _ := strconv.ParseInt(s[1], 10, 64)

		cards := parseHand(s[0])
		handLevel := getHandValue(cards)

		hands[i] = hand{bid: int(bid), cards: cards, handLevel: handLevel}
	}

	return hands
}

func getHandValue(cards [5]rank) level {
	valueCount := make(map[rank]int, 0)

	for _, v := range cards {
		valueCount[v] = valueCount[v] + 1
	}

	if len(valueCount) == 1 {
		return fiveKind
	} else if len(valueCount) == 2 {
		// full house or four kind
		if v := valueCount[cards[0]]; v == 1 || v == 4 {
			return fourKind
		} else {
			return fullHouse
		}
	} else if len(valueCount) == 3 {
		// threekind or 2 pair
		var m int
		for _, v := range valueCount {
			if v > m {
				m = v
			}
		}

		if m == 3 {
			return threeKind
		} else {
			return twoPair
		}
	} else if len(valueCount) == 4 {
		return onePair
	} else {
		return high
	}
}

func parseHand(line string) [5]rank {
	var ranks [5]rank

	for i, v := range line {
		if v >= '0' && v <= '9' {
			ranks[i] = rank(v - '0' - 1)
		} else if v == 'T' {
			ranks[i] = rank(9)
		} else if v == 'J' {
			ranks[i] = rank(10)
		} else if v == 'Q' {
			ranks[i] = rank(11)
		} else if v == 'K' {
			ranks[i] = rank(12)
		} else if v == 'A' {
			ranks[i] = rank(13)
		} else {
			panic(fmt.Sprintf("Invalid parse: %c", v))
		}
	}
	return ranks
}
