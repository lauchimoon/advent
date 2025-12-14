package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
    "strconv"
)

const FILEPATH = "resources/001input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    var x int64 = 50
    zeroCount := 0

    fmt.Println(x)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        rot := line[0]
        rotVal, err := strconv.ParseInt(line[1:], 10, 64)
        if err != nil {
            log.Fatalf("failed to parse integer")
        }

        if rot == 'R' {
            x = (x + rotVal) % 100
        } else {
            x = (x - rotVal) % 100
            if x < 0 {
                x += 100
            }
        }

        if x == 0 {
            zeroCount++
        }
    }

    fmt.Println(zeroCount)
}
