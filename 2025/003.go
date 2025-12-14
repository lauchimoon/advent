package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
)

const FILEPATH = "resources/003input.txt"

func main() {
    banks := LoadBanks()
    sum := 0

    for _, bank := range banks {
        m := FindMax(bank)
        sum += m
    }

    fmt.Println(sum)
}

func LoadBanks() []string {
    banks := []string{}
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        banks = append(banks, scanner.Text())
    }

    return banks
}

func FindMax(s string) int {
    lenS := len(s)
    m := -1
    for l := 0; l < lenS; l++ {
        for r := l + 1; r < lenS; r++ {
            var n int = int(10*(s[l] - '0') + (s[r] - '0'))
            if n > m {
                m = n
            }
        }
    }

    return m
}
