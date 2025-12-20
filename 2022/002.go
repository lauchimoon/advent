package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

const (
    FILEPATH = "./resources/002input.txt"

    ROCK_OPP    = "A"
    PAPER_OPP   = "B"
    SCISSOR_OPP = "C"

    ROCK_PLAYER    = "X"
    PAPER_PLAYER   = "Y"
    SCISSOR_PLAYER = "Z"

)

var (
    shapePoints = map[string]int{
        "X": 1, "Y": 2, "Z": 3,
    }
)

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    total := 0
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        plays := strings.Fields(scanner.Text())
        opponent := plays[0]
        player := plays[1]
        score := CalculateScore(opponent, player)
        total += score
    }
    fmt.Println(total)
}

func CalculateScore(opponent, player string) int {
    score := 0

    if opponent == ROCK_OPP {
        switch player {
            case PAPER_PLAYER:
                score = 6
                break
            case ROCK_PLAYER:
                score = 3
                break
            case SCISSOR_PLAYER:
                score = 0
                break
        }
    } else if opponent == PAPER_OPP {
        switch player {
            case PAPER_PLAYER:
                score = 3
                break
            case ROCK_PLAYER:
                score = 0
                break
            case SCISSOR_PLAYER:
                score = 6
                break
        }
    } else if opponent == SCISSOR_OPP {
        switch player {
            case PAPER_PLAYER:
                score = 0
                break
            case ROCK_PLAYER:
                score = 6
                break
            case SCISSOR_PLAYER:
                score = 3
                break
        }
    }

    return score + shapePoints[player]
}
