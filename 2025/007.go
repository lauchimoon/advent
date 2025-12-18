package main

import (
    "bufio"
    //"io"
    "fmt"
    "log"
    "os"
    "strings"
)

const (
    FILEPATH = "./resources/007input.txt"
    LEN = 141
)

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("couldn't open file '%s'", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Scan()
    paths := make([]int, LEN)
    i := strings.Index(scanner.Text(), "S")
    total := 0
    paths[i] = 1

    for scanner.Scan() {
        s := scanner.Text()
        for i := range s {
            if s[i] != '^' {
                continue
            }
            if paths[i] > 0 {
                total++
            }

            paths[i-1] += paths[i]
            paths[i+1] += paths[i]
            paths[i] = 0
        }
    }
    fmt.Println(total)
}
