package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

const (
    FILEPATH = "./resources/011input.txt"

    STATE_FLOOR = -1
    STATE_FREE = 0
    STATE_OCC = 1
)

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
    seats := MakeSeats(lines)
    for !StatesEqual(seats, NextState(seats)) {
        seats = NextState(seats)
    }

    return CountOccupied(seats)
}

func MakeSeats(lines []string) [][]int {
    seats := make([][]int, len(lines))
    for i := range seats {
        seats[i] = make([]int, len(lines[i]))
    }

    for y := range lines {
        for x := range lines[y] {
            seats[y][x] = GetSeatState(lines[y][x])
        }
    }

    return seats
}

func GetSeatState(seat byte) int {
    if seat == 'L' {
        return STATE_FREE
    } else if seat == '#' {
        return STATE_OCC
    }

    return STATE_FLOOR
}

func StatesEqual(state1, state2 [][]int) bool {
    if len(state1) != len(state2) && len(state1[0]) != len(state2[0]) {
        return false
    }

    for y := range state1 {
        for x := range state1[y] {
            if state1[y][x] != state2[y][x] {
                return false
            }
        }
    }

    return true
}

func NextState(state [][]int) [][]int {
    newState := make([][]int, len(state))
    for i := range newState {
        newState[i] = make([]int, len(state[i]))
    }

    for y := range state {
        for x := range state[y] {
            switch state[y][x] {
            case STATE_FREE:
                newState[y][x] = UpdateFreeSeat(state, y, x)
                break
            case STATE_OCC:
                newState[y][x] = UpdateOccupiedSeat(state, y, x)
                break
            case STATE_FLOOR:
                newState[y][x] = STATE_FLOOR
                break
            }
        }
    }
    return newState
}

func UpdateFreeSeat(state [][]int, y, x int) int {
    neighbours := GetNeighbours(state, y, x)
    for _, n := range neighbours {
        if n == STATE_OCC {
            return STATE_FREE
        }
    }
    return STATE_OCC
}

func UpdateOccupiedSeat(state [][]int, y, x int) int {
    neighbours := GetNeighbours(state, y, x)
    occupiedCount := 0

    for _, n := range neighbours {
        if n == STATE_OCC {
            occupiedCount++
        }
    }

    if occupiedCount > 4 {
        return STATE_FREE
    }

    return STATE_OCC
}

func GetNeighbours(state [][]int, y, x int) []int {
    neighbours := []int{}
    directions := []int{-1, 0, 1}
    for _, dy := range directions {
        for _, dx := range directions {
            if y + dy >= 0 && y + dy < len(state) && x + dx >= 0 && x + dx < len(state[0]) {
                neighbours = append(neighbours, state[y + dy][x + dx])
            }
        }
    }

    return neighbours
}

func CountOccupied(state [][]int) int {
    occupiedCount := 0
    for y := range state {
        for x := range state[y] {
            if state[y][x] == STATE_OCC {
                occupiedCount++
            }
        }
    }
    return occupiedCount
}
