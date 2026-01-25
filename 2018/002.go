package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    f, _ := os.Open("resources/002input.txt")
    defer f.Close()

    lines := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    fmt.Println("Part 1:", Part1(lines))
    fmt.Println("Part 2:", Part2(lines))
}

func Part1(lines []string) int {
    twoLetter := 0
    threeLetter := 0
    for _, line := range lines {
        freqMap := map[rune]int{}
        count(line, &freqMap)
        updateValues(&twoLetter, &threeLetter, freqMap)
    }
    return twoLetter*threeLetter
}

func count(line string, freqMap *map[rune]int) {
    for _, c := range line {
        (*freqMap)[c]++
    }
}

func updateValues(twoLetter, threeLetter *int, freqMap map[rune]int) {
    checkedTwo := false
    checkedThree := false
    for _, v := range freqMap {
        if v == 2 && !checkedTwo {
            (*twoLetter)++
            checkedTwo = true
        } else if v == 3 && !checkedThree {
            (*threeLetter)++
            checkedThree = true
        }
    }
}

func Part2(lines []string) string {
    length := len(lines[0])
    for i := 0; i < length; i++ {
        seenMasks := map[string]bool{}
        for _, line := range lines {
            mask := line[:i] + line[i+1:]
            if seenMasks[mask] {
                return mask
            }
            seenMasks[mask] = true
        }
    }
    return ":("
}
