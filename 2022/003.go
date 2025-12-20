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
        compartment := scanner.Text()
        mid := len(compartment)/2
        firstHalf := compartment[:mid]
        secondHalf := compartment[mid:]
        commonElement := FindCommonType(firstHalf, secondHalf)
        priority := GetElementPriority(commonElement)
        total += priority
    }
    fmt.Println(total)
}

func FindCommonType(firstHalf, secondHalf string) rune {
    for _, cfh := range firstHalf {
        for _, csh := range secondHalf {
            if csh == cfh {
                return csh
            }
        }
    }
    return '?'
}

func GetElementPriority(element rune) int {
    if unicode.IsLower(element) {
        return int(element - 'a' + 1)
    } else {
        return int(element - 'A' + 1) + 26
    }
}
