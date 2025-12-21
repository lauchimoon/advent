package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

const FILEPATH = "./resources/005input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    seats := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        seats = append(seats, scanner.Text())
    }

    fmt.Printf("Part 1: %d\n", Part1(seats))
}

func Part1(seats []string) int {
    maxSeatId := 0
    for _, seat := range seats {
        seatId := GetRow(seat)*8 + GetCol(seat)
        if seatId > maxSeatId {
            maxSeatId = seatId
        }
    }
    return maxSeatId
}

func GetRow(seat string) int {
    low, high := 0, 127
    for _, c := range seat[:7] {
        if c == 'F' {
            high = (high + low)/2
        } else {
            low = (high + low)/2 + 1
        }
    }
    return low
}

func GetCol(seat string) int {
    low, high := 0, 7
    for _, c := range seat[7:] {
        if c == 'R' {
            low = (high + low)/2 + 1
        } else {
            high = (high + low)/2
        }
    }
    return low
}
