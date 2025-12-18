package main

import (
    "bufio"
    "io"
    "fmt"
    "log"
    "strconv"
    "strings"
    "sort"
    "os"
)

const (
    FILEPATH = "./resources/005input.txt"
    RANGES_END = 185
    NUMBERS_BEGIN = RANGES_END + 1
)

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("couldn't open file '%s'", FILEPATH)
    }
    defer f.Close()

    fileString := LoadFileString(f)
    ranges := LoadRanges(fileString[:RANGES_END])

    fmt.Printf("Part 1: %v\n", Part1(fileString, ranges))
    fmt.Printf("Part 2: %v\n", Part2(ranges))
}

func LoadFileString(f io.Reader) []string {
    s := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        s = append(s, scanner.Text())
    }
    return s
}

func LoadRanges(fileString []string) []uint64 {
    ranges := []uint64{}
    for _, s := range fileString {
        numbersString := strings.Split(s, "-")
        n1, _ := strconv.ParseUint(numbersString[0], 10, 64)
        n2, _ := strconv.ParseUint(numbersString[1], 10, 64)
        ranges = append(ranges, n1, n2)
    }
    return ranges
}

func Part1(fileString []string, ranges []uint64) int {
    ids := LoadIDs(fileString[NUMBERS_BEGIN:])
    count := 0
    lenRanges := len(ranges)
    for id := range ids {
        for i := 0; i < lenRanges-1; i += 2 {
            if id >= ranges[i] && id <= ranges[i+1] && !ids[id] {
                ids[id] = true
                count++
            }
        }
    }
    return count
}

// Sweep-line algorithm
func Part2(ranges []uint64) uint64 {
    type Event struct {
        Pos  uint64
        Type int
    }
    events := []Event{}
    lenRanges := len(ranges)
    for i := 0; i < lenRanges-1; i += 2 {
        events = append(events, Event{ranges[i], 1})
        events = append(events, Event{ranges[i+1], -1})
    }

    sort.Slice(events, func(i, j int) bool {
        if events[i].Pos == events[j].Pos {
            return events[i].Type > events[j].Type
        }
        return events[i].Pos < events[j].Pos
    })

    var count uint64 = 0
    openIntervals := 0
    var lastPos uint64 = 0

    for _, event := range events {
        if openIntervals > 0 {
            count += event.Pos - lastPos
        }

        if openIntervals == 0 && event.Type == 1 {
            count++
        }

        openIntervals += event.Type
        lastPos = event.Pos
    }

    return count
}

func LoadIDs(fileString []string) map[uint64]bool {
    ids := map[uint64]bool{}
    for _, s := range fileString {
        n, _ := strconv.ParseUint(s, 10, 64)
        ids[n] = false
    }
    return ids
}
