package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
)

const day = 10
const runInput = true

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

	for _, v := range potentialAdjacent {
		if v.x < 0 || v.x >= len(lines) || v.y < 0 || v.y >= len(lines[0]) {
			continue
		}
		m, n := getAdjacentNodes(v, lines)
		// check and see if the thing connects to start
		if m != startingPosition && n != startingPosition {
			continue
		}
		ret := dfs(v, startingPosition, &lines, nil)

		if ret > 0 {
			return (ret + 1) / 2
		}
	}
	return -1
}

func dfs(curr point, parent point, lines *[]string, colors *[][]bool) int {

	if curr.x < 0 || curr.x >= len(*lines) || curr.y < 0 || curr.y >= len((*lines)[0]) {
		return -1e9
	}
	if (*lines)[curr.x][curr.y] == 'S' {
		return 0
	}

	first, second := getAdjacentNodes(curr, *lines)

	if colors != nil {
		// mark it as visited
		markVisited(curr, curr, colors)
	}
	if first != parent {
		markVisited(curr, first, colors)
		return 1 + dfs(first, curr, lines, colors)
	}
	if second != parent {
		markVisited(curr, second, colors)
		return 1 + dfs(second, curr, lines, colors)
	}
	log.Println("something went wrong", curr, first, second, rune((*lines)[curr.x][curr.y]))
	return -1e8
}

func markVisited(p, t point, colors *[][]bool) {
	if colors == nil {
		return
	}

	(*colors)[p.x*2][p.y*2] = true
	if t != p {
		dx := t.x - p.x
		dy := t.y - p.y
		(*colors)[p.x*2+dx][p.y*2+dy] = true
	}
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

func solvePart2(lines []string) int {
	// let's brainstorm some ideas here
	// one thing that would be worth doing is having some sort of "visited" array, to know the bounds of the loop
	// then, we could iterate over the entire graph and check if each piece is bound by the loop
	// how do we tell if a piece is in the loop?
	// let's think of this as a series of steps, and consider the problem as a graph coloring
	// 1. Find and mark every single node that is part of the loop. let's say as a 1
	// 2. "color" the graph by checking if each node is colored. if its not colored,
	// 		this probably wont work, remember the funky case where things can squeeze between

	// an entirely different approach could be to PAD THE ENTIRE array by doubling its size
	// now, each odd index is actually representing if a gap is blocked
	// we can fill these in since we know what things are in the main loop
	// a | would block north and south gaps, but the gaps on its left and right remain open (since things could squeeze between)
	// this sounds like the most promising approach I'm having to determine bounds.

	// after a long walk, shopoping, and a shower i have a better idea
	// firstly, create an array called colors where everything is inited to 0
	// colors should be be size 2n - 1, since it will represent the gaps in between pipes
	// for each element of the loop, color it with a 1
	// then, iter over the OUTSIDE EDGES of the color array
	// start dfs from each spot to see if it gets blocked by the loop
	// after iterating over every edge, we iterate back over the original array
	// check to see if each node is colored. every uncolored node should be counted
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

	potentialAdjacent := []point{
		{x: startingPosition.x, y: startingPosition.y - 1},
		{x: startingPosition.x, y: startingPosition.y + 1},
		{x: startingPosition.x + 1, y: startingPosition.y},
		{x: startingPosition.x - 1, y: startingPosition.y},
	}

	colors := make([][]bool, 2*len(lines)-1)
	for i := range colors {
		colors[i] = make([]bool, -1+2*len(lines[0]))
	}

	markVisited(startingPosition, startingPosition, &colors)
	var direction [4]bool

	for i, v := range potentialAdjacent {
		if v.x < 0 || v.x >= len(lines) || v.y < 0 || v.y >= len(lines[0]) {
			continue
		}
		m, n := getAdjacentNodes(v, lines)
		// check and see if the thing connects to start
		if m != startingPosition && n != startingPosition {
			continue
		}
		markVisited(v, m, &colors)
		markVisited(v, n, &colors)
		ret := dfs(v, startingPosition, &lines, &colors)

		if ret > 0 {
			// we have to figure out which shape it should be
			direction[i] = true
		}
	}

	for i := range lines {

		if !colors[2*i][0] {
			dfsColor(point{x: 2 * i, y: 0}, &colors)
		}
		if !colors[2*i][len(colors[0])-1] {
			dfsColor(point{x: 2 * i, y: len(colors[0]) - 1}, &colors)
		}
	}

	for i := range lines[0] {
		if !colors[len(colors)-1][2*i] {
			dfsColor(point{x: len(colors) - 1, y: 2 * i}, &colors)
		}
		if !colors[0][2*i] {
			dfsColor(point{x: 0, y: 2 * i}, &colors)
		}
	}
	// now we know where the loop is (direction wise)

	unVisited := 0

	for i, r := range lines {
		for j := range r {
			if !colors[2*i][2*j] {
				log.Println(i, j)
				unVisited++
			}
		}
	}
	return unVisited
}

func dfsColor(p point, visited *[][]bool) {
	if p.x < 0 || p.x >= len(*visited) || p.y < 0 || p.y >= len((*visited)[0]) {
		return
	}
	if (*visited)[p.x][p.y] {
		return
	}
	(*visited)[p.x][p.y] = true

	dfsColor(point{x: p.x + 1, y: p.y}, visited)
	dfsColor(point{x: p.x - 1, y: p.y}, visited)
	dfsColor(point{x: p.x, y: p.y + 1}, visited)
	dfsColor(point{x: p.x, y: p.y - 1}, visited)
}
