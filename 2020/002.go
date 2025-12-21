package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

const FILEPATH = "./resources/002input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    passwords := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        passwords = append(passwords, scanner.Text())
    }

    fmt.Printf("Part 1: %d\n", Part1(passwords))
    fmt.Printf("Part 2: %d\n", Part2(passwords))
}

func Part1(passwords []string) int {
    total := 0
    for _, passwd := range passwords {
        if IsValidPassword(passwd, 1) {
            total++
        }
    }
    return total
}

func IsValidPassword(passwd string, part int) bool {
    info := strings.Split(passwd, ": ")
    policy := strings.Fields(info[0])
    low, high := GetLetterTimes(policy[0])
    key := info[1]
    letter := policy[1][0]

    if part == 1 {
        letterCount := GetLetterCount(key, letter)
        return letterCount >= low && letterCount <= high
    } else {
        return (key[low - 1] == letter && key[high - 1] != letter) ||
               (key[high - 1] == letter && key[low - 1] != letter)
    }
}

func GetLetterTimes(letterRange string) (int, int) {
    split := strings.Split(letterRange, "-")
    low, _ := strconv.Atoi(split[0])
    high, _ := strconv.Atoi(split[1])
    return low, high
}

func GetLetterCount(key string, letter byte) int {
    count := 0
    for _, c := range key {
        if byte(c) == letter {
            count++
        }
    }

    return count
}

func Part2(passwords []string) int {
    total := 0
    for _, passwd := range passwords {
        if IsValidPassword(passwd, 2) {
            total++
        }
    }
    return total
}
