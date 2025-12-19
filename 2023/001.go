package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "unicode"

    "github.com/dlclark/regexp2"
)

const FILEPATH = "./resources/001input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    lines := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    fmt.Printf("Part 1: %d\n", Part1(lines))
    fmt.Printf("Part 2: %d\n", Part2(lines))
}

func Part1(lines []string) int {
    total := 0
    for _, line := range lines {
        src := line
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
    return total
}

func Part2(lines []string) int {
    total := 0
    numberWords := strings.Fields("one two three four five six seven eight nine")
    pattern := "(?=(" + strings.Join(numberWords, "|") + "|[0-9]))"
    re := regexp2.MustCompile(pattern, 0)
    replace := map[string]int{
        "one": 1, "two": 2, "three": 3,
        "four": 4, "five": 5, "six": 6,
        "seven": 7, "eight": 8, "nine": 9,
        "1": 1, "2": 2, "3": 3,
        "4": 4, "5": 5, "6": 6,
        "7": 7, "8": 8, "9": 9,
    }

    for _, line := range lines {
        numbers := []string{}
        match, _ := re.FindStringMatch(line)
        for match != nil {
            numbers = append(numbers, match.GroupByNumber(1).String())
            match, _ = re.FindNextMatch(match)
        }

        first := replace[numbers[0]]
        second := replace[numbers[len(numbers)-1]]
        total += 10*first + second
    }

    return total
}
