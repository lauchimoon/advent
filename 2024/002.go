package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

const FILEPATH = "./resources/002input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    lines := []string{}
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    fmt.Printf("Part 1: %d\n", Part1(lines))
    fmt.Printf("Part 2: %d\n", Part2(lines))
}

func Part1(lines []string) int {
    count := 0
    for _, line := range lines {
        if IsSafe(SequenceToSlice(line)) {
            count++
        }
    }
    return count
}

func SequenceToSlice(seq string) []int64 {
    split := strings.Split(seq, " ")
    list := []int64{}
    for _, s := range split {
        x, _ := strconv.ParseInt(s, 10, 64)
        list = append(list, x)
    }
    return list
}

func IsSafe(list []int64) bool {
    return (AllInc(list) || AllDec(list)) && CorrectDiff(list)
}

func AllInc(list []int64) bool {
    lenList := len(list)
    for i := 0; i < lenList-1; i++ {
        if list[i] >= list[i+1] {
            return false
        }
    }
    return true
}

func AllDec(list []int64) bool {
    lenList := len(list)
    for i := 0; i < lenList-1; i++ {
        if list[i] <= list[i+1] {
            return false
        }
    }
    return true
}

func CorrectDiff(list []int64) bool {
    lenList := len(list)
    for i := 0; i < lenList-1; i++ {
        diff := Abs(list[i]-list[i+1])
        if diff < 1 || diff > 3 {
            return false
        }
    }
    return true
}

func Abs(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x
}

func Part2(lines []string) int {
    total := 0
    for _, s := range lines {
        l := SequenceToSlice(s)
        idx := 0
        for idx < len(l) {
            if IsSafe(SliceDelete(l, idx)) {
                total++
                break
            }
            idx++
        }
    }
    return total
}

func SliceDelete(s []int64, i int) []int64 {
    sNew := []int64{}
    for idx := range s {
        if idx != i {
            sNew = append(sNew, s[idx])
        }
    }
    return sNew
}
