package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func SolveWith[T any, V any](day int, solvePart1 func([]T) V, solvePart2 func([]T) V, parse func([]string) []T) {
	filePath := fmt.Sprintf("day%d/sample.txt", day)
	lines := ReadTextFromFileToStringSlice(filePath)
	d := parse(lines)
	log.Println("Part 1: ", solvePart1(d))
	log.Println("Part 2: ", solvePart2(d))
}
