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

    answers := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        answers = append(answers, scanner.Text())
    }

    fmt.Printf("Part 1: %d\n", Part1(answers))
    fmt.Printf("Part 2: %d\n", Part2(answers))
}

func Part1(answers []string) int {
    answersString := ""
    total := 0
    for _, answer := range answers {
        answersString += answer
        if answer == "" {
            total += GetNumberAnswers(answersString)
            answersString = ""
        }
    }

    total += GetNumberAnswers(answersString)
    return total
}

func GetNumberAnswers(answers string) int {
    questionsAnswered := map[rune]bool{}
    for _, c := range answers {
        questionsAnswered[c] = true
    }

    return len(questionsAnswered)
}

func Part2(answers []string) int {
    groups := GetGroups(answers)
    total := 0

    for _, g := range groups {
        total += GetNumberAllAnswered(g)
    }

    return total
}

func GetGroups(answers []string) [][]string {
    filtered := []string{}
    for _, answer := range answers {
        if answer == "" {
            filtered = append(filtered, answer)
        }
    }

    groups := make([][]string, len(filtered) + 1)
    i := 0

    for _, answer := range answers {
        if answer != "" {
            groups[i] = append(groups[i], answer)
        } else {
            i++
        }
    }

    return groups
}

func GetNumberAllAnswered(group []string) int {
    count := map[rune]int{}
    lenGroup := len(group)
    total := 0

    for _, g := range group {
        for _, c := range g {
            count[c]++
        }
    }

    for _, v := range count {
        if v == lenGroup {
            total++
        }
    }

    return total
}
