package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"
)

const FILEPATH = "./resources/001input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    calories := []int64{}
    scanner := bufio.NewScanner(f)
    total := int64(0)
    for scanner.Scan() {
        text := scanner.Text()

        if text != "" {
            cal, _ := strconv.ParseInt(scanner.Text(), 10, 64)
            total += cal
        } else {
            calories = append(calories, total)
            total = 0
        }
    }

    // Make sure to append last value
    calories = append(calories, total)

    sort.Slice(calories, func(i, j int) bool {
        return calories[i] > calories[j]
    })

    sumTopThree := calories[0] + calories[1] + calories[2]
    fmt.Printf("Part 1: %v\n", Max(calories))
    fmt.Printf("Part 2: %v\n", sumTopThree)
}

func Max(slice []int64) int64 {
    m := int64(0)
    for _, x := range slice {
        if x > m {
            m = x
        }
    }

    return m
}
