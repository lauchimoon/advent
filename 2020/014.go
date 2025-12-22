package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "regexp"
)

const FILEPATH = "./resources/014input.txt"

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

    fmt.Printf("Part 1: %v\n", Part1(lines))
}

func Part1(lines []string) int64 {
    values := map[int64]int64{}
    mask := ""
    for _, line := range lines {
        fields := strings.Split(line, " = ")

        if fields[0] == "mask" {
            mask = fields[1]
        } else {
            re := regexp.MustCompile(`\d+`)
            memAddr, _ := strconv.ParseInt(re.FindString(fields[0]), 10, 64)
            value, _ := strconv.ParseInt(fields[1], 10, 64)
            values[memAddr] = ApplyMask(value, mask)
        }
    }

    total := int64(0)
    for _, v := range values {
        total += v
    }

    return total
}

func ApplyMask(n int64, mask string) int64 {
    mask0, _ := strconv.ParseInt(strings.ReplaceAll(mask, "X", "1"), 2, 64)
    mask1, _ := strconv.ParseInt(strings.ReplaceAll(mask, "X", "0"), 2, 64)
    return (n & mask0) | mask1
}

func MaskToInt(mask string) int64 {
    mask = strings.ReplaceAll(mask, "X", "1")
    maskInt, _ := strconv.ParseInt(mask, 2, 64)
    return maskInt
}
