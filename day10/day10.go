package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
)

const day = 10
const runInput = !false

func main() {
	var filePath string
	if runInput {
		filePath = fmt.Sprintf("day%d/input.txt", day)
	} else {
		filePath = fmt.Sprintf("day%d/sample.txt", day)
	}

	lines := utils.ReadTextFromFileToStringSlice(filePath)

	log.Println("Part 1: ", solvePart1(lines))
	log.Println("Part 2: ", solvePart2(lines))
}

type point struct {
	x int
	y int
}

func solvePart1(lines []string) int {
	var startingPosition = func() point {
		for i := range lines {
			for j := range lines[i] {
				if lines[i][j] == 'S' {
					return point{x: int(i), y: int(j)}
				}
			}
		}
		return point{x: -1, y: -1}
	}()

	distances := make([][]int, len(lines))
	for i := range distances {
		distances[i] = make([]int, len(lines[i]))
	}

	// run a bfs
	// starting at start position, find all connected adjacent pipes
	// is it possible for the pipe to be connected on every side?
	// no, we have to figure out which type the thing should be.
	// how can we do that? we'd basically have to run a bfs and see if it ever returns to start.
	// first thing to do is figure out which adjacents are part of thel oop

	// okay, so for each adjacent pipe that might be connected:
	// start a regular bfs from it. if the queue becomes empty AND we haven't reached the starting position, then none of those are part of the loop and should be marked accordingly
	// otherwise, great, we found the distances
	// maybe on nonloops, we should just reset the distances array in between.

	// start from an adjacent, not starting position

	// better idea: since there is only one loop, we are actually doing a dfs.
	// each 'node' has 2 neighbors it connects to in the loop
	// really, we know the furthest position is just going to be sizeOfLoop/2
	// we could either fill in a distances array as we go and return back the backwards count to get each node its accurate count
	// alt if a node starts returning and it didn't reach s, then we have some sentinel value where we start marking the distances as negative
	// since its a loop, if on exploration A we find the loop, then when we check the alt loop B, it will already have a distance
	// thus we use that trip to find the max distance
	// or, if we never readjust on the way back, then its even easier: when we find node B, which connects to S and has a positive value, we have its distance since we passed it down
	// this means that we would know the path length, and can thus transform it to determine what the furthest point in the loop must be.
	// its close enough that can be modified pretty easily if we need distandce of every point later

	// alternatively - we could try setting S to each possible pipe then running the BFS search
	// the trick would be, if a node is enqueued twice, its done

	// i prefer the older approach of just calculating the length of the path (DFS) until we hopefully run into S again

	potentialAdjacent := []point{
		{x: startingPosition.x, y: startingPosition.y - 1},
		{x: startingPosition.x, y: startingPosition.y + 1},
		{x: startingPosition.x + 1, y: startingPosition.y},
		{x: startingPosition.x - 1, y: startingPosition.y},
	}

	log.Println(startingPosition)
	for _, v := range potentialAdjacent {
		if v.x < 0 || v.x >= len(distances) || v.y < 0 || v.y >= len(distances[0]) {
			continue
		}
		m, n := getAdjacentNodes(v, lines)
		// check and see if the thing connects to start
		if m != startingPosition && n != startingPosition {
			continue
		}
		ret := dfs(v, startingPosition, &lines)

		if ret > 0 {
			return (ret + 1) / 2
		}
	}
	return -1
}

func dfs(curr point, parent point, lines *[]string) int {

	if curr.x < 0 || curr.x >= len(*lines) || curr.y < 0 || curr.y >= len((*lines)[0]) {
		return -1e9
	}
	if (*lines)[curr.x][curr.y] == 'S' {
		return 0
	}

	first, second := getAdjacentNodes(curr, *lines)

	if first != parent {
		return 1 + dfs(first, curr, lines)
	}
	if second != parent {
		return 1 + dfs(second, curr, lines)
	}
	log.Println("something went wrong", curr, first, second, rune((*lines)[curr.x][curr.y]))
	return -1e8
}

func getAdjacentNodes(curr point, lines []string) (point, point) {
	switch lines[curr.x][curr.y] {
	case '|':
		return point{x: curr.x - 1, y: curr.y}, point{x: curr.x + 1, y: curr.y}
	case '-':
		return point{x: curr.x, y: curr.y - 1}, point{x: curr.x, y: curr.y + 1}
	case 'L':
		return point{x: curr.x - 1, y: curr.y}, point{x: curr.x, y: curr.y + 1}
	case 'J':
		return point{x: curr.x - 1, y: curr.y}, point{x: curr.x, y: curr.y - 1}
	case '7':
		return point{x: curr.x + 1, y: curr.y}, point{x: curr.x, y: curr.y - 1}
	case 'F':
		return point{x: curr.x + 1, y: curr.y}, point{x: curr.x, y: curr.y + 1}
	default:
		return curr, curr
	}
}

type queue []point

func (q *queue) enqueue(x point) {
	*q = append(*q, x)
}

func (q *queue) dequeue() point {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}
func solvePart2(lines []string) int {
	return 1
}
