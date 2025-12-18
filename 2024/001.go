package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
    "strconv"
)

const FILEPATH = "resources/001input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    leftList := []int64{}
    rightList := []int64{}

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        split := strings.Split(scanner.Text(), " ")
        lhs, _ := strconv.ParseInt(split[0], 10, 64)
        rhs, _ := strconv.ParseInt(split[len(split)-1], 10, 64)

        leftList = append(leftList, lhs)
        rightList = append(rightList, rhs)
    }

    sort.Slice(leftList, func(i, j int) bool { return leftList[i] < leftList[j] })
    sort.Slice(rightList, func(i, j int) bool { return rightList[i] < rightList[j] })

    // Both lists have the same length
    lenList := len(leftList)
    var total int64 = 0
    for i := 0; i < lenList; i++ {
        total += Abs(leftList[i] - rightList[i])
    }
    fmt.Println(total)
}

func Abs(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x
}
