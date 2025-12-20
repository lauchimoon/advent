package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
    "strconv"
)

const FILEPATH = "./resources/006input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    stats := LoadStats(f)
    times := []int64{}
    distances := []int64{}

    for _, stat := range strings.Fields(stats[0])[1:] {
        x, _ := strconv.ParseInt(stat, 10, 64)
        times = append(times, x)
    }
    for _, stat := range strings.Fields(stats[1])[1:] {
        x, _ := strconv.ParseInt(stat, 10, 64)
        distances = append(distances, x)
    }

    fmt.Printf("Part 1: %d\n", Part1(times, distances))
    fmt.Printf("Part 2: %d\n", Part2(times, distances))
}

func LoadStats(f io.Reader) []string {
    stats := []string{}
    scanner := bufio.NewScanner(f)

    scanner.Scan()
    stats = append(stats, scanner.Text())

    scanner.Scan()
    stats = append(stats, scanner.Text())

    return stats
}

func Part1(times, distances []int64) int {
    total := 1
    nRaces := len(times)

    for i := 0; i < nRaces; i++ {
        top := times[i]
        waysToWin := 0
        for buttonTime := int64(0); buttonTime <= top; buttonTime++ {
            movement := buttonTime*(times[i] - buttonTime)
            if movement > distances[i] {
                waysToWin++
            }
        }

        total *= waysToWin
    }
    return total
}

func Part2(times, distances []int64) int {
    trueTime := ConcatenateSlice(times)
    trueDistance := ConcatenateSlice(distances)

    waysToWin := 0
    for buttonTime := int64(0); buttonTime <= trueTime; buttonTime++ {
        movement := buttonTime*(trueTime - buttonTime)
        if movement > trueDistance {
            waysToWin++
        }
    }

    return waysToWin
}

func ConcatenateSlice(slice []int64) int64 {
    sx := ""
    for _, x := range slice {
        sx += strconv.FormatInt(x, 10)
    }
    
    v, _ := strconv.ParseInt(sx, 10, 64)
    return v
}
