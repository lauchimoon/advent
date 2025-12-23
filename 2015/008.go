package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

const FILEPATH = "./resources/008input.txt"

func main() {
	f, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatalf("failed to open file %s", FILEPATH)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}

func Part1(lines []string) int {
	total := 0
	for _, line := range lines {
		total += len(line) - CalculateLength(line)
	}
	return total
}

func CalculateLength(s string) int {
	trueLen := 0
	for i := 1; i < len(s) - 1; i++ {
		if s[i] == '\\' {
			if s[i+1] == 'x' {
				i += 3
			} else {
				i++
			}
		}
		trueLen++
	}

	return trueLen
}

func Part2(lines []string) int {
	total := 0
	for _, line := range lines {
		total += CalculateLengthNew(line) - len(line)
	}
	return total
}

func CalculateLengthNew(s string) int {
	newLen := 2
	for i := 0; i < len(s); i++ {
		if s[i] == '"' {
			newLen += 2
		} else if s[i] == '\\' {
			newLen += 2
		} else {
			newLen++
		}
	}

	return newLen
}
