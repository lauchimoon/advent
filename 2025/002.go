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

    var total int64 = 0
    scanner := bufio.NewScanner(f)
    scanner.Scan()
    ranges := strings.Split(scanner.Text(), ",")
    for _, rs := range ranges {
        r := BuildRange(rs)
        for _, i := range r {
            if IsInvalid(i) {
                total += i
            }
        }
    }
    fmt.Println(total)
}

func BuildRange(rangeString string) []int64 {
    rang := []int64{}
    split := strings.Split(rangeString, "-")
    start, _ := strconv.ParseInt(split[0], 10, 64)
    end, _ := strconv.ParseInt(split[1], 10, 64)

    for i := start; i <= end; i++ {
        rang = append(rang, i)
    }
    return rang
}

func IsInvalid(x int64) bool {
    sx := strconv.FormatInt(x, 10)
    lenSx := len(sx)
    if lenSx % 2 != 0 {
        return false
    }

    mid := lenSx/2
    for i := 0; i < mid; i++ {
        if sx[i] != sx[i+mid] {
            return false
        }
    }

    return true
}
