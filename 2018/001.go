package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    f, _ := os.Open("resources/001input.txt")
    defer f.Close()

    lines := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    fmt.Println("Part 1:", Part1(lines))
    fmt.Println("Part 2:", Part2(lines))
}

func Part1(lines []string) int {
    freq := 0
    for _, line := range lines {
        val, _ := strconv.Atoi(line)
        freq += val
    }
    return freq
}

func Part2(lines []string) int {
    seenFreq := map[int]bool{}
    freq := 0
    lineIdx := 0
    for true {
        val, _ := strconv.Atoi(lines[lineIdx])
        freq += val
        if seenFreq[freq] {
            return freq
        }
        seenFreq[freq] = true
        lineIdx = (lineIdx + 1)%len(lines)
    }
    return -1
}
