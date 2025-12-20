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

    total := 1
    nRaces := len(times)

    for i := 0; i < nRaces; i++ {
        top := times[i]
        waysToWin := 0
        for buttonTime := int64(0); buttonTime <= top; buttonTime++ {
            movement := buttonTime*(times[i] - buttonTime)
            fmt.Printf("hold %v: %vmm\n", buttonTime, movement)
            if movement > distances[i] {
                waysToWin++
            }
        }

        total *= waysToWin
    }
    fmt.Println(total)
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
