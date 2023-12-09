package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strings"
)

const day = 9
const runInput = true

type int int64

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

func solvePart1(sequences []sequence) int {
	var ret int = 0
	for _, s := range sequences {
		o := calculateNextValue(s)
		ret += o
	}
	return ret
}

func solvePart2(sequences []sequence) int {
	var ret int = 0
	for _, s := range sequences {
		nums := make([]int, len(s.nums))

		for i, v := range s.nums {
			nums[len(nums)-1-i] = v
		}

		o := calculateNextValue(sequence{nums})
		ret += o
	}
	return ret
}
func calculateNextValue(s sequence) int {
	// recursively, return the next value using the differences
	// method: if i am all 0, return 0
	// else, return the recursive call + my last item
	if finished(s) {
		return 0
	}

	diffs := calculateDifferences(s)

	return calculateNextValue(diffs) + s.nums[len(s.nums)-1]
}
func finished(s sequence) bool {
	for _, v := range s.nums {
		if v != 0 {
			return false
		}
	}
	return true
}
func calculateDifferences(s sequence) sequence {
	nums := s.nums
	out := make([]int, len(nums)-1)

	for i := 1; i < len(nums); i++ {
		out[i-1] = nums[i] - nums[i-1]
	}

	return sequence{nums: out}

}

type sequence struct {
	nums []int
}

func parse(lines []string) []sequence {
	out := make([]sequence, len(lines))

	for i, v := range lines {
		s := strings.Fields(v)

		l := make([]int, len(s))

		for j, p := range s {
			l[j] = int(utils.Atoi(p))
		}

		out[i] = sequence{l}
	}

	return out
}
