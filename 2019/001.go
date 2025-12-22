package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

const FILEPATH = "./resources/001input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    masses := []int{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        x, _ := strconv.Atoi(scanner.Text())
        masses = append(masses, x)
    }

	fmt.Printf("Part 1: %d\n", Part1(masses))
	fmt.Printf("Part 2: %d\n", Part2(masses))
}

func Part1(masses []int) int {
	total := 0
	for _, mass := range masses {
		total += mass/3 - 2
	}
	return total
}

func Part2(masses []int) int {
	total := 0
	for _, mass := range masses {
		x := mass/3 - 2
		total += x
		for x > 0 {
			x = x/3 - 2
			if x <= 0 {
				break
			}
			total += x
		}
	}
	return total
}
