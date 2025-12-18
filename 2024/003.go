package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "unicode"
)

const FILEPATH = "./resources/003input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    total := 0
    for scanner.Scan() {
        total += ParseMul(scanner.Text())
    }

    fmt.Println(total)
}

func ParseMul(s string) int {
    openParen := false
    lhs := 0
    rhs := 0
    total := 0

    for i := 0; i < len(s); i++ {
        if s[i] == 'm' && s[i+1] == 'u' && s[i+2] == 'l' {
            i += 3
            if s[i] == '(' {
                openParen = true
                i++
            }

            if unicode.IsDigit(rune(s[i])) {
                for unicode.IsDigit(rune(s[i])) {
                    lhs = 10*lhs + int(s[i] - '0')
                    i++
                }
            }

            if s[i] == ',' {
                i++
            }

            if unicode.IsDigit(rune(s[i])) {
                for unicode.IsDigit(rune(s[i])) {
                    rhs = 10*rhs + int(s[i] - '0')
                    i++
                }
            }

            if openParen && s[i] == ')' {
                openParen = false
                total += lhs*rhs
            }

            lhs = 0
            rhs = 0
        }
    }

    return total
}
