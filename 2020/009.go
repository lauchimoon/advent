package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

const FILEPATH = "./resources/009input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()
    numbers := []int64{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        x, _ := strconv.ParseInt(scanner.Text(), 10, 64)
        numbers = append(numbers, x)
    }

    fmt.Printf("Part 1: %v\n", Part1(numbers))
}

func Part1(numbers []int64) int64 {
    offset := 25
    i := offset
    for IsSum(numbers[i-offset:i], numbers[i]) {
        i++
    }

    return numbers[i]
}

func IsSum(numbers []int64, x int64) bool {
    for _, i := range numbers {
        for _, j := range numbers {
            if i != j && i + j == x {
                return true
            }
        }
    }

    return false
}
