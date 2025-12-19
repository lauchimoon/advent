package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "unicode"
)

const FILEPATH = "./resources/001input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    total := 0
    for scanner.Scan() {
        src := scanner.Text()
        L := 0
        R := len(src)-1
        first, second := ' ', ' '

        for !unicode.IsDigit(first) || !unicode.IsDigit(second) {
            first = rune(src[L])
            if !unicode.IsDigit(first) {
                L++
            }

            second = rune(src[R])
            if !unicode.IsDigit(second) {
                R--
            }
        }

        total += int(10*(first - '0') + (second - '0'))
    }
    fmt.Println(total)
}
