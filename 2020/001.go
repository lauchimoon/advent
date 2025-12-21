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

    for _, x := range numbers {
        for _, y := range numbers {
            if x + y == 2020 {
                fmt.Println(x*y)
                return
            }
        }
    }
}
