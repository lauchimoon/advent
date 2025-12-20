package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

const FILEPATH = "./resources/004input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    total := 0
    for scanner.Scan() {
        pair := strings.Split(scanner.Text(), ",")
        range1 := MakeRange(pair[0])
        range2 := MakeRange(pair[1])

        if IsContained(range1, range2) {
            total++
        }
    }
    fmt.Println(total)
}

func MakeRange(rangeStr string) []int64 {
    r := make([]int64, 2)
    split := strings.Split(rangeStr, "-")
    r[0], _ = strconv.ParseInt(split[0], 10, 64)
    r[1], _ = strconv.ParseInt(split[1], 10, 64)
    return r
}

func IsContained(range1 []int64, range2 []int64) bool {
    leftmost1, leftmost2 := range1[0], range2[0]
    rightmost1, rightmost2 := range1[1], range2[1]

    return (leftmost2 >= leftmost1 && rightmost2 <= rightmost1) ||
           (leftmost1 >= leftmost2 && rightmost1 <= rightmost2)
}
