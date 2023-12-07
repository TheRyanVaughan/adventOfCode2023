package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"math"
	"strings"
)

const runInput = true
const day = 6

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
	log.Println("Part 2: ", solvePart2(lines))
}

type race struct {
	time     int64
	distance int64
}

func parse(lines []string) []race {
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	var races []race
	for i := range times {
		races = append(races, race{time: utils.Atoi(strings.TrimSpace(times[i])), distance: utils.Atoi(strings.TrimSpace(distances[i]))})
	}
	return races
}

func solvePart1(races []race) int64 {
	var ret int64

	ret = 1
	for _, r := range races {
		first, second := quadraticEquation(-1, r.time, -1*r.distance)

		low := int64(math.Floor(1.0 + first))
		high := int64(math.Ceil(second - 1.0))

		ret *= high - low + 1
	}
	// find every possible way to finish the race in the time
	// can spend x milliseconds to get x speed
	// obviously greedy
	// you don't move until you're done choosing your speed
	// if speed is x, then we win the race as long as x * (time-x) > dist
	// -x^2 + xtime - dist > 0
	// solve for both intercepts

	return ret
}

func solvePart2(lines []string) int64 {
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]

	time := utils.Atoi(strings.Join(times, ""))
	distance := utils.Atoi(strings.Join(distances, ""))

	log.Println(time, distance)
	first, second := quadraticEquation(-1, time, -1*distance)

	low := int64(math.Floor(1.0 + first))
	high := int64(math.Ceil(second - 1.0))

	ret := high - low + 1
	return ret

}

func quadraticEquation(a int64, b int64, c int64) (float64, float64) {
	// returns an array of floats, each root
	h := float64(a)
	k := float64(b)
	l := float64(c)
	x := ((k * -1.0) + math.Sqrt((k*k)-(4*h*l))) / (2 * h)
	y := ((k * -1.0) - math.Sqrt((k*k)-(4*h*l))) / (2 * h)
	return x, y
}
