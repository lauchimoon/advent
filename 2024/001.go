package main

import (
    "bufio"
    "fmt"
    "io"
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

    leftList, rightList := LoadLists(f)

    fmt.Printf("Part 1: %v\n", Part1(leftList, rightList))
    fmt.Printf("Part 2: %v\n", Part2(leftList, rightList))
}

func LoadLists(f io.Reader) ([]int64, []int64) {
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
    return leftList, rightList
}

func Part1(leftList, rightList []int64) int64 {
    sort.Slice(leftList, func(i, j int) bool { return leftList[i] < leftList[j] })
    sort.Slice(rightList, func(i, j int) bool { return rightList[i] < rightList[j] })

    // Both lists have the same length
    lenList := len(leftList)
    var total int64 = 0
    for i := 0; i < lenList; i++ {
        total += Abs(leftList[i] - rightList[i])
    }
    return total
}

func Abs(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x
}

func Part2(leftList,  rightList []int64) int64 {
    var total int64 = 0
    for _, number := range leftList {
        times := CountFromSlice(rightList, number)
        total += times*number
    }
    return total
}

func CountFromSlice(slice []int64, n int64) int64 {
    var count int64 = 0
    for _, x := range slice {
        if x == n {
            count++
        }
    }
    return count
}
