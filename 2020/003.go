package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

const FILEPATH = "./resources/003input.txt"

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
    fmt.Printf("Part 2: %d\n", Part2(lines))
}

func Part1(lines []string) int {
    return CountTrees(lines, 1, 3)
}

func CountTrees(lines []string, dy, dx int) int {
    width := len(lines[0])
    height := len(lines)
    total := 0
    y := dy
    x := dx % width

    for y < height {
        if lines[y][x] == '#' {
            total++
        }

        y += dy
        x = (x + dx) % width
    }

    return total
}

func Part2(lines []string) int {
    slopes := [][]int{
        {1, 1},
        {1, 3},
        {1, 5},
        {1, 7},
        {2, 1},
    }

    total := 1
    for _, slope := range slopes {
        total *= CountTrees(lines, slope[0], slope[1])
    }

    return total
}
