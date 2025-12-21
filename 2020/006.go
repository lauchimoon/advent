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
