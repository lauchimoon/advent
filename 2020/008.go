package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

type Instruction struct {
    Operation string
    Argument  int
    Visited   bool
}

const FILEPATH = "./resources/008input.txt"

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
}

func Part1(lines []string) int {
    program := LoadProgram(lines)
    insIndex := 0
    accumulator := 0
    for insIndex < len(program) && !program[insIndex].Visited {
        ParseInstruction(&program[insIndex], &accumulator, &insIndex)
        insIndex++
    }

    return accumulator
}

func LoadProgram(lines []string) []Instruction {
    program := []Instruction{}
    for _, line := range lines {
        ins := strings.Fields(line)

        op := ins[0]
        arg, _ := strconv.Atoi(ins[1])
        program = append(program, Instruction{op, arg, false})
    }

    return program
}

func ParseInstruction(ins *Instruction, accumulator, insIndex *int) {
    ins.Visited = true
    switch ins.Operation {
    case "nop":
        break
    case "acc":
        (*accumulator) += ins.Argument
        break
    case "jmp":
        (*insIndex) += ins.Argument - 1
        break
    default:
        break
    }
}
