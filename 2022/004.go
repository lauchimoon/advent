package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

const FILEPATH = "./resources/004input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    pairs := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        pairs = append(pairs, scanner.Text())
    }

    fmt.Printf("Part 1: %d\n", Part1(pairs))
    fmt.Printf("Part 2: %d\n", Part2(pairs))
}

func Part1(pairs []string) int {
    total := 0
    for _, line := range pairs {
        pair := strings.Split(line, ",")
        range1 := MakeRange(pair[0])
        range2 := MakeRange(pair[1])

        if IsContained(range1, range2) {
            total++
        }
    }
    return total
}

func MakeRange(rangeStr string) []int64 {
    r := make([]int64, 2)
    split := strings.Split(rangeStr, "-")
    r[0], _ = strconv.ParseInt(split[0], 10, 64)
    r[1], _ = strconv.ParseInt(split[1], 10, 64)
    return r
}

func IsContained(range1 []int64, range2 []int64) bool {
    x1, x2 := range1[0], range2[0]
    y1, y2 := range1[1], range2[1]

    return (x2 >= x1 && y2 <= y1) ||
           (x1 >= x2 && y1 <= y2)
}

func Part2(pairs []string) int {
    total := 0
    for _, line := range pairs {
        pair := strings.Split(line, ",")
        range1 := MakeRange(pair[0])
        range2 := MakeRange(pair[1])

        if RangesOverlap(range1, range2) {
            total++
        }
    }
    return total
}

func RangesOverlap(range1, range2 []int64) bool {
    x1, x2 := range1[0], range2[0]
    y1, y2 := range1[1], range2[1]

    return max(x1, x2) - min(y1, y2) <= 0
}
