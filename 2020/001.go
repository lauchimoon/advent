package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

const FILEPATH = "./resources/001input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()
    numbers := []int{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        x, _ := strconv.Atoi(scanner.Text())
        numbers = append(numbers, x)
    }

    fmt.Printf("Part 1: %d\n", Part1(numbers))
    fmt.Printf("Part 2: %d\n", Part2(numbers))
}

func Part1(numbers []int) int {
    for _, x := range numbers {
        for _, y := range numbers {
            if x + y == 2020 {
                return x*y
            }
        }
    }
    return -1
}

func Part2(numbers []int) int {
    for _, x := range numbers {
        for _, y := range numbers {
            for _, z := range numbers {
                if x + y + z == 2020 {
                    return x*y*z
                }
            }
        }
    }
    return -1
}
