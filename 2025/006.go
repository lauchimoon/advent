package main

import (
    "bufio"
    "io"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

const FILEPATH = "./resources/006input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("couldn't open file '%s'", FILEPATH)
    }
    defer f.Close()

    rows := LoadRows(f)
    total := CalculateTotal(rows)
    fmt.Println(total)
}

func LoadRows(f io.Reader) [][]string {
    rows := [][]string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        rows = append(rows, SplitRow(scanner.Text()))
    }
    return rows
}

func SplitRow(text string) []string {
    split := []string{}
    splitText := strings.Split(text, " ")
    for _, s := range splitText {
        if s != "" {
            split = append(split, s)
        }
    }
    return split
}

func CalculateTotal(rows [][]string) uint64 {
    nColumns := len(rows[0])
    var total uint64 = 0

    for i := 0; i < nColumns; i++ {
        numbers := []int{}
        op := ""
        for row := range rows {
            s := rows[row][i]
            num, _ := strconv.Atoi(s)

            if s == "+" || s == "*" {
                op = s
            } else {
                numbers = append(numbers, num)
            }
        }

        total += ApplyOp(numbers, op)
    }
    return total
}

func ApplyOp(numbers []int, op string) uint64 {
    var result uint64 = 0
    if op == "*" {
        result++
    }

    for _, num := range numbers {
        if op == "+" {
            result += uint64(num)
        } else {
            result *= uint64(num)
        }
    }
    return result
}
