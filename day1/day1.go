package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const inputFileName = "input1.txt"

/*
Ways I'd improve this:
 1. My slicing in GetDigit means I create a new string eveyrtime. Use a "has prefix" function to check if the substring starting at pos[i] is a word string\
 2. solvepart 1 is a bit rough. Perhaps I could generalize GetDigit to accept an array of strings that map to their value (pref a dict)
    and then pass in an empty one for part 1, so then the rest of the logic would be the same as part two
*/
func main() {

	file, err := os.Open(inputFileName)

	if err != nil {
		panic("Could not read the contents of the file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))

}

func solvePart1(lines []string) int {
	var sum int
	for _, line := range lines {
		var first, last int
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if first == 0 {
					first = atoi(c)
				}
				last = atoi(c)
			}
		}
		sum += 10*first + last
	}
	return sum
}

func solvePart2(lines []string) int {
	var sum int
	for _, line := range lines {
		sum += lineValue(line)
	}
	return sum
}

func lineValue(line string) int {
	var first, last int
	for i, _ := range line {
		val, err := getDigit(line, i)
		if err == nil {
			if first == 0 {
				first = val
			}
			last = val
		}
	}

	return 10*first + last
}

func atoi(c rune) int {
	return int(c - '0')
}

var numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getDigit(line string, pos int) (int, error) {
	if line[pos] >= '0' && line[pos] <= '9' {
		return atoi(rune(line[pos])), nil
	}

	endIdx := len(line)
	for n, num := range numbers {
		l := len(num)
		if line[pos:min(pos+l, endIdx)] == num {
			return n, nil
		}
	}
	return 0, errors.New("Not a digit")
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
