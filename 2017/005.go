package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"slices"
)

const FILEPATH = "./resources/005input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

	scanner := bufio.NewScanner(f)
	jumps := []int{}
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		jumps = append(jumps, x)
	}

	jumpsOriginal := slices.Clone(jumps)
	fmt.Println("Part 1:", Solve(jumps, 1))
	jumps = slices.Clone(jumpsOriginal)
	fmt.Println("Part 2:", Solve(jumps, 2))
}

func Solve(jumps []int, part int) int {
	idx := 0
	count := 0
	for idx < len(jumps) {
		prevOffset := jumps[idx]
		increase := 1
		if part == 2 && prevOffset >= 3 {
			increase = -1
		}

		jumps[idx] += increase
		idx += prevOffset
		count++
	}

	return count
}
