package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

const FILEPATH = "./resources/015input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Scan()
    labels := strings.Split(scanner.Text(), ",")
    total := 0
    for _, s := range labels {
        total += HASH(s)
    }
    fmt.Println(total)
}

func HASH(s string) int {
    h := 0
    for _, c := range s {
        h += int(c)
        h *= 17
        h %= 256
    }
    return h
}

