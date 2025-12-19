package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

const FILEPATH = "./resources/004input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    cards := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        cards = append(cards, scanner.Text())
    }

    total := 0
    for _, card := range cards {
        cardsInfo := strings.Split(card, " | ")
        winnerCards := ConstructWinner(strings.Fields(strings.Split(cardsInfo[0], ": ")[1]))
        myCards := strings.Fields(cardsInfo[1])

        firstFound := false
        points := 0

        for _, c := range myCards {
            _, isWinner := winnerCards[c]
            if isWinner {
                if !firstFound {
                    firstFound = true
                    points++
                } else {
                    points *= 2
                }
            }
        }
        total += points
    }
    fmt.Println(total)
}

func ConstructWinner(cards []string) map[string]bool {
    winnerCards := map[string]bool{}
    for _, c := range cards {
        winnerCards[c] = true
    }

    return winnerCards
}
