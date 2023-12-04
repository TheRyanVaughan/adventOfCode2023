package utils

import (
	"bufio"
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
