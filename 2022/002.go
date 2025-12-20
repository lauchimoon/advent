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
        plays := strings.Fields(line)
        opponent := plays[0]
        player := plays[1]
        score := CalculateScore(opponent, player)
        total += score
    }
    return total
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

func Part2(lines []string) int {
    total := 0
    for _, line := range lines {
        plays := strings.Fields(line)
        opponent := plays[0]
        player := DeterminePlayer(opponent, plays[1])
        score := CalculateScore(opponent, player)
        total += score
    }
    return total
}

func DeterminePlayer(opponent, end string) string {
    if opponent == ROCK_OPP {
        switch end {
        case "X": return SCISSOR_PLAYER
        case "Y": return ROCK_PLAYER
        case "Z": return PAPER_PLAYER
        default: return "?"
        }
    } else if opponent == PAPER_OPP {
        switch end {
        case "X": return ROCK_PLAYER
        case "Y": return PAPER_PLAYER
        case "Z": return SCISSOR_PLAYER
        default: return "?"
        }
    } else if opponent == SCISSOR_OPP {
        switch end {
        case "X": return PAPER_PLAYER
        case "Y": return SCISSOR_PLAYER
        case "Z": return ROCK_PLAYER
        default: return "?"
        }
    }

    return "?"
}
