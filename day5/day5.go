package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const runInput = !false
const day = 5

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

func parse(lines []string) input {
	var out input
	s := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seeds := make([]int64, len(s))

	for i := range s {
		seeds[i], _ = strconv.ParseInt(s[i], 10, 64)
	}
	out.seeds = seeds

	var currMap mapContents
	for i := 2; i < len(lines); i++ {
		v := lines[i]
		if strings.Contains(v, "map") {
			currMap.name = strings.Split(v, " map")[0]
		} else if v == "" {
			out.maps = append(out.maps, currMap)
			currMap = mapContents{}
		} else {
			ints := strings.Split(v, " ")
			var d description
			d.dest, _ = strconv.ParseInt(ints[0], 10, 64)
			d.src, _ = strconv.ParseInt(ints[1], 10, 64)
			d.len, _ = strconv.ParseInt(ints[2], 10, 64)
			currMap.info = append(currMap.info, d)
		}
	}
	out.maps = append(out.maps, currMap)

	return out
}

type input struct {
	seeds []int64
	maps  []mapContents
}
type mapContents struct {
	name string
	info []description
}

type description struct {
	dest int64
	src  int64
	len  int64
}

func solvePart1(lines input) int64 {

	seedMap := make(map[int64]int64, len(lines.seeds))
	for _, s := range lines.seeds {
		seedMap[s] = s
	}

	for _, m := range lines.maps {
		newMap := make(map[int64]int64)
		for _, d := range m.info {
			for k, s := range seedMap {
				diff := s - d.src
				if diff >= 0 && diff < d.len {
					newMap[k] = d.dest + diff
					log.Println(newMap[k], newMap[k]-d.dest+d.src, s)
				}
			}
		}

		for k, v := range newMap {
			seedMap[k] = v
		}

	}

	minKey := lines.seeds[0]
	for k, v := range seedMap {
		if v < seedMap[minKey] {
			minKey = k
		}
	}

	return seedMap[minKey]
}

func solvePart2(lines input) int64 {
	var minKey int64 = 1
	maps := lines.maps

	for {
		// find first time a location maps to any range of seed
		currVal := minKey
		for i := len(maps) - 1; i >= 0; i-- {
			for _, r := range maps[i].info {
				// i'm in the right range
				diff := currVal - r.dest
				if diff >= 0 && diff < r.len {
					// log.Println(minKey, i, currVal, r)
					currVal = r.src + diff
					break
				}
			}
		}

		// log.Println(currVal)
		// find if this is a valid seed
		for i := 0; i < len(lines.seeds); i += 2 {
			if currVal >= lines.seeds[i] && currVal < lines.seeds[i]+lines.seeds[i+1] {
				return minKey
			}
		}
		minKey++
	}
	// how can we map a valid destination to its previous src?
	// consider arbitrary map x, that maps a -> b via the following: given a, aRangeStart, bRangeStart, we get b as follows:
	// a - aRangeStart := diff, which should be the same offset of b from bRangeStart
	// conversely, the inverse mapping b -> a should work identically
	// b - bRangeStart := diff, which means adding diff to aRangeStart ought to give a.

	return minKey
}
