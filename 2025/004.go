package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
)

const (
    FILEPATH = "resources/004input.txt"
    SIZE_X = 135
    SIZE_Y = 135
)

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("couldn't open file '%s'", FILEPATH)
    }
    defer f.Close()

    rollGrid := LoadRollGrid(f)
    count := 0

    for y := 0; y < SIZE_Y; y++ {
        for x := 0; x < SIZE_X; x++ {
            neighbours := CountNeighbours(rollGrid, y, x)
            if neighbours < 4 && rollGrid[y][x] == '@' {
                count++
            }
        }
    }

    fmt.Println(count)
}

func LoadRollGrid(f io.Reader) [SIZE_Y][SIZE_X]rune {
    row := 0
    col := 0
    rollGrid := [SIZE_Y][SIZE_X]rune{}

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        for _, c := range line {
            rollGrid[row][col] = c
            col++
        }
        row++
        col = 0
    }

    return rollGrid
}

func CountNeighbours(grid [SIZE_Y][SIZE_X]rune, y, x int) int {
    count := 0
    for yy := -1; yy <= 1; yy++ {
        for xx := -1; xx <= 1; xx++ {
            isRoll := CheckCell(grid, y, x, yy, xx)
            if isRoll {
                count++
            }
        }
    }
    return count
}

func CheckCell(grid [SIZE_Y][SIZE_X]rune, y, x, yy, xx int) bool {
    if (xx == 0 && yy == 0) {
        return false
    }

    if (y + yy < 0 || y + yy >= SIZE_Y || x + xx < 0 || x + xx >= SIZE_X) {
        return false
    }

    return grid[y][x] == '@' && grid[y+yy][x+xx] == '@'
}
