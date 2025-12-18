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
    count := 0
    for scanner.Scan() {
        if IsSafe(scanner.Text()) {
            count++
        }
    }
    fmt.Println(count)
}

func IsSafe(seq string) bool {
    split := strings.Split(seq, " ")
    list := []int64{}
    for _, s := range split {
        x, _ := strconv.ParseInt(s, 10, 64)
        list = append(list, x)
    }

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
