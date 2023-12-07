package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadTextFromFileToStringSlice(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("error was not nil: ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func SolveWith[T any, V any](day int, solvePart1 func([]T) V, solvePart2 func([]T) V, parse func([]string) []T, runInput bool) {
	var filePath string
	if runInput {
		filePath = fmt.Sprintf("day%d/input.txt", day)
	} else {
		filePath = fmt.Sprintf("day%d/sample.txt", day)
	}

	lines := ReadTextFromFileToStringSlice(filePath)
	d := parse(lines)
	log.Println("Part 1: ", solvePart1(d))
	log.Println("Part 2: ", solvePart2(d))
}

func NoParse(input []string) []string {
	return input
}

func Atoi(s string) int64 {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return x
}
