package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "slices"
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
    fmt.Printf("Part 2: %d\n", Part2(seats))
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

func Part2(seats []string) int {
    ids := []int{}
    for _, seat := range seats {
        seatId := GetRow(seat)*8 + GetCol(seat)
        ids = append(ids, seatId)
    }

    for row := 0; row < 128; row++ {
        for col := 0; col < 8; col++ {
            id := row*8 + col
            if slices.Index(ids, id) == -1 {
                // Calculated based on graphic representation
                // of free/unexistent seats
                if id > 74 {
                    return id
                }
            }
        }
    }
    return -1
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
