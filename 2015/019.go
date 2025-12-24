package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
)

const FILEPATH = "./resources/019input.txt"

func main() {
	f, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatalf("failed to open file %s", FILEPATH)
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", Part1(lines))
}

func Part1(lines []string) int {
	replacements := LoadReplacements(lines)
	start := lines[len(lines)-1]
	molecules := map[string]bool{}

	for replaceTarget, replace := range replacements {
		for _, replacement := range replace {
			i := 0
			sub := ReplaceNth(start, replaceTarget, replacement, i)
			for sub != start {
				molecules[sub] = true
				i++
				sub = ReplaceNth(start, replaceTarget, replacement, i)
			}
		}
	}
	return len(molecules)
}

func LoadReplacements(lines []string) map[string][]string {
	replacements := map[string][]string{}

	for _, line := range lines {
		if line == "" {
			break
		}

		split := strings.Split(line, " => ")
		src := split[0]
		replacement := split[1]
		replacements[src] = append(replacements[src], replacement)
	}

	return replacements
}

func ReplaceNth(src, target, replace string, n int) string {
	if n < 0 || target == "" {
		return src
	}

	currentIdx := -1
	occurrenceCount := 0
	for i := 0; i <= n; i++ {
		offset := currentIdx + 1
		if currentIdx == -1 {
			offset = 0
		}

		found := strings.Index(src[offset:], target)
		if found == -1 {
			return src
		}

		currentIdx = offset + found
		occurrenceCount++
	}

	return src[:currentIdx] + replace + src[currentIdx + len(target):]
}
