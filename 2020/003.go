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
}

func Part1(lines []string) int {
    width := len(lines[0])
    height := len(lines)
    dy := 1
    dx := 3
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
