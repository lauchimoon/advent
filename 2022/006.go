package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

const FILEPATH = "./resources/006input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Scan()
    idx := FindMarkerIndex(scanner.Text())
    fmt.Println(idx + 4)
}

func FindMarkerIndex(stream string) int {
    for i := 0; i < len(stream)-4; i++ {
        if AllDifferent(stream[i:i+4]) {
               return i
        }
    }

    return -1
}

func AllDifferent(s string) bool {
    m := map[rune]bool{}
    for _, c := range s {
        if m[c] {
            return false
        }

        m[c] = true
    }

    return true
}
