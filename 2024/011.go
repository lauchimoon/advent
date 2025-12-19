package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

const FILEPATH = "./resources/011input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Scan()
    stones := LoadStones(scanner.Text())

    fmt.Printf("Part 1: %v\n", Blink(stones, 25))
    fmt.Printf("Part 2: %v\n", Blink(stones, 75))
}

func LoadStones(seq string) map[int64]int64 {
    stones := map[int64]int64{}
    for _, s := range strings.Fields(seq) {
        x, _ := strconv.ParseInt(s, 10, 64)
        stones[x]++
    }

    return stones
}

func Blink(stones map[int64]int64, n int) int64 {
    var total int64 = 0
    for i := 0; i < n; i++ {
        newStones := map[int64]int64{}
        for val, count := range stones {
            for _, x := range UpdateStone(val) {
                newStones[x] += count
            }
        }
        stones = newStones
    }

    for _, count := range stones {
        total += count
    }
    return total
}

func UpdateStone(x int64) []int64 {
    if x == 0 {
        return []int64{1}
    }

    sx := strconv.FormatInt(x, 10)
    lenSx := len(sx)
    if lenSx % 2 == 0 {
        mid := lenSx/2
        left, _ := strconv.ParseInt(sx[:mid], 10, 64)
        right, _ := strconv.ParseInt(sx[mid:], 10, 64)
        return []int64{left, right}
    }

    return []int64{2024*x}
}
