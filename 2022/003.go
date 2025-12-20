package main

import (
    "bufio"
    "fmt"
    "log"
    "maps"
    "os"
    "slices"
    "unicode"
)

const FILEPATH = "./resources/003input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    rucksacks := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        rucksacks = append(rucksacks, scanner.Text())
    }

    fmt.Printf("Part 1: %d\n", Part1(rucksacks))
    fmt.Printf("Part 2: %d\n", Part2(rucksacks))
}

func Part1(rucksacks []string) int {
    total := 0
    for _, compartment := range rucksacks {
        mid := len(compartment)/2
        firstHalf := compartment[:mid]
        secondHalf := compartment[mid:]
        commonElement := FindCommonType(firstHalf, secondHalf)
        priority := GetElementPriority(commonElement)
        total += priority
    }
    return total
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

func Part2(rucksacks []string) int {
    total := 0
    i := 0
    for i < len(rucksacks)/3 {
        commonElement := FindCommonThree(rucksacks[3*i:(3*i + 3)])
        priority := GetElementPriority(commonElement)
        total += priority
        i++
    }

    return total
}

func FindCommonThree(rucksacks []string) rune {
    first := rucksacks[0]
    second := rucksacks[1]
    third := rucksacks[2]

    commonMap := map[rune]bool{}
    for _, c := range first {
        commonMap[c] = true
    }

    commonMap = Intersect(commonMap, second)
    commonMap = Intersect(commonMap, third)

    return slices.Collect(maps.Keys(commonMap))[0]
}

func Intersect(common map[rune]bool, s string) map[rune]bool {
    newCommon := map[rune]bool{}
    for _, c := range s {
        if common[c] {
            newCommon[c] = true
        }
    }
    return newCommon
}
