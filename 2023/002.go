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

    games := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        games = append(games, scanner.Text())
    }

    fmt.Printf("Part 1: %d\n", Part1(games))
    fmt.Printf("Part 2: %d\n", Part2(games))
}

func Part1(games []string) int {
    possible := true
    total := 0
    id := 1
    for _, game := range games {
        possible = true
        info := strings.Split(game, ": ")
        sets := strings.Split(info[1], "; ")

        var r, g, b int64 = 0, 0, 0

        for _, set := range sets {
            cubes := strings.Split(set, ", ")
            for _, cube := range cubes {
                cubeInformation := strings.Fields(cube)
                amount, _ := strconv.ParseInt(cubeInformation[0], 10, 64)
                color := cubeInformation[1]

                r, g, b = UpdateRGB(color, amount, r, g, b)

                if r > 12 || g > 13 || b > 14 {
                    possible = false
                    break
                }
            }

            if !possible {
                break
            }
        }

        if possible {
            total += id
        }
        id++
    }
    return total
}

func UpdateRGB(color string, amount int64, r, g, b int64) (int64, int64, int64) {
    switch color {
    case "red":
        r = amount
    case "green":
        g = amount
    case "blue":
        b = amount
    }

    return r, g, b
}

func Part2(games []string) int64 {
    var total int64 = 0
    var mr, mg, mb int64 = 0, 0, 0
    for _, game := range games {
        info := strings.Split(game, ": ")
        sets := strings.Split(info[1], "; ")

        var r, g, b int64 = 0, 0, 0

        for _, set := range sets {
            cubes := strings.Split(set, ", ")
            for _, cube := range cubes {
                cubeInformation := strings.Fields(cube)
                amount, _ := strconv.ParseInt(cubeInformation[0], 10, 64)
                color := cubeInformation[1]

                r, g, b = UpdateRGB(color, amount, r, g, b)
                mr, mg, mb = UpdateMax(mr, mg, mb, r, g, b)
            }
        }

        total += mr*mg*mb
        mr, mg, mb = 0, 0, 0
    }
    return total
}

func UpdateMax(mr, mg, mb, r, g, b int64) (int64, int64, int64) {
    if r > mr {
        mr = r
    }
    if g > mg {
        mg = g
    }
    if b > mb {
        mb = b
    }

    return mr, mg, mb
}
