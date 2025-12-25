package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sort"
)

const FILEPATH = "./resources/004input.txt"

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

	fmt.Println("Part 1:", Solve(lines, 1))
	fmt.Println("Part 2:", Solve(lines, 2))
}

func Solve(lines []string, part int) int {
	total := 0
	check := false
	for _, line := range lines {
		if part == 1 {
			check = CheckNoDupWords(line)
		} else {
			check = CheckNoAnagrams(line)
		}

		if check {
			total++
		}
	}

	return total
}

func CheckNoDupWords(line string) bool {
	words := map[string]bool{}
	for _, word := range strings.Fields(line) {
		if words[word] {
			return false
		}
		words[word] = true
	}
	return true
}

func CheckNoAnagrams(line string) bool {
	words := map[string]bool{}
	for _, word := range strings.Fields(line) {
		sortedWord := SortString(word)
		if words[sortedWord] {
			return false
		}
		words[sortedWord] = true
	}

	return true
}

func SortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
