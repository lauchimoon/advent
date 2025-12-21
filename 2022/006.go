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
    fmt.Printf("Part 1: %d\n", FindMarkerIndex(scanner.Text(), 4) + 4)
    fmt.Printf("Part 2: %d\n", FindMarkerIndex(scanner.Text(), 14) + 14)
}

func FindMarkerIndex(stream string, distinct int) int {
    for i := 0; i < len(stream)-distinct; i++ {
        if AllDifferent(stream[i:i+distinct]) {
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
