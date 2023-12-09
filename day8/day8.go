package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strings"
)

const day = 8
const runInput = true

func main() {
	var filePath string
	if runInput {
		filePath = fmt.Sprintf("day%d/input.txt", day)
	} else {
		filePath = fmt.Sprintf("day%d/sample.txt", day)
	}

	lines := utils.ReadTextFromFileToStringSlice(filePath)
	d := parse(lines)

	// log.Println("Part 1: ", solvePart1(d))
	log.Println("Part 2: ", solvePart2(d))
}

func solvePart1(in input) int {
	curr := "AAA"

	steps := 0

	currInputIndex := 0
	for curr != "ZZZ" {
		curr = in.nodeMap[curr][in.instructions[currInputIndex]]
		currInputIndex = (currInputIndex + 1) % len(in.instructions)
		steps++
	}

	return steps
}

// this only works because the cycles are nice in the input
// i.e, no starting node will cross multiple exists
func solvePart2(in input) int64 {
	var currs []string

	for k := range in.nodeMap {
		if k[2] == 'A' {
			currs = append(currs, k)
		}
	}

	steps := 0
	currInputIndex := 0

	cycleMap := make(map[string]int)

	finished := func() bool {
		if len(cycleMap) != len(currs) {
			return false
		}

		for _, v := range cycleMap {
			if v <= 0 {
				return false
			}
		}
		return true
	}

	// map each key to its cycle period. end once all keys are filled in
	for !finished() {
		for i, v := range currs {
			currs[i] = in.nodeMap[v][in.instructions[currInputIndex]]
			if currs[i][2] == 'Z' {
				if cycleMap[currs[i]] == 0 {
					cycleMap[currs[i]] = -1 * steps
				} else if cycleMap[currs[i]] < 0 {
					cycleMap[currs[i]] = steps - (-1 * cycleMap[currs[i]])
				}
			}
		}
		currInputIndex = (currInputIndex + 1) % len(in.instructions)
		steps++
	}

	var lcm int64 = 1

	for _, v := range cycleMap {
		product := lcm * int64(v)
		gcd := func(a, b int64) int64 {
			if a < b {
				t := b
				b = a
				a = t
			}

			for a%b != 0 {
				r := a % b
				a = b
				b = r
			}
			return b
		}
		lcm = product / gcd(lcm, int64(v))
	}

	return lcm
}

type direction int

const (
	left direction = iota
	right
)

type input struct {
	instructions []direction
	nodeMap      map[string][2]string
}

func parse(lines []string) input {
	directions := lines[0]
	instructions := make([]direction, len(directions))
	for i, v := range directions {
		if v == 'L' {
			instructions[i] = left
		} else {
			instructions[i] = right
		}
	}

	nodes := make(map[string][2]string)
	for _, v := range lines {
		if strings.Contains(v, "=") {
			srcNode := v[:3]

			tupleSplit := strings.Split(v, ",")
			leftChoice := tupleSplit[0][len(tupleSplit[0])-3:]
			rightChoice := tupleSplit[1][1:4]

			t := [2]string{leftChoice, rightChoice}

			nodes[srcNode] = t
		}
	}

	return input{instructions: instructions, nodeMap: nodes}
}
