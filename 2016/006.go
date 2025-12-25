package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	FILEPATH = "./resources/006input.txt"
	MESSAGE_LEN = 8
)

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

	scanner := bufio.NewScanner(f)
	messages := []string{}
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}

	counts := CountLetters(messages)
	fmt.Println("Part 1:", GetErrorCorrectedMessage(counts, 1))
	fmt.Println("Part 2:", GetErrorCorrectedMessage(counts, 2))
}

func CountLetters(messages []string) []map[rune]int {
	counts := make([]map[rune]int, MESSAGE_LEN)
	for i := range counts {
		// Make space for all letters
		counts[i] = make(map[rune]int, 26)
	}

	for _, message := range messages {
		for i, c := range message {
			counts[i][c]++
		}
	}

	return counts
}

func GetErrorCorrectedMessage(counts []map[rune]int, part int) string {
	errorCorrected := make([]rune, MESSAGE_LEN)
	for i := range counts {
		var char rune
		if part == 1 {
			char = FindMostFrequent(counts[i])
		} else {
			char = FindLeastFrequent(counts[i])
		}

		errorCorrected[i] = char
	}
	return string(errorCorrected)
}

func FindMostFrequent(count map[rune]int) rune {
	maxCount := 0
	var char rune

	for k, v := range count {
		if v > maxCount {
			maxCount = v
			char = k
		}
	}
	return char
}

func FindLeastFrequent(count map[rune]int) rune {
	minCount := math.MaxInt32
	var char rune

	for k, v := range count {
		if v < minCount {
			minCount = v
			char = k
		}
	}
	return char
}
