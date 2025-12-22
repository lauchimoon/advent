package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "slices"
)

const (
	FILEPATH = "./resources/002input.txt"

	INSTRUCTION_ADD = 1
	INSTRUCTION_MUL = 2
	INSTRUCTION_HALT = 99
)

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
	scanner.Scan()
	program := Map(strings.Split(scanner.Text(), ","), func (s string) int {
        x, _ := strconv.Atoi(s); return x
	})

    initialProgram := slices.Clone[[]int, int](program)
    fmt.Printf("Part 1: %d\n", Part1(program, 12, 2))
    fmt.Printf("Part 2: %d\n", Part2(initialProgram, program))
}

func Map(slice []string, f func(string) int) []int {
	newSlice := []int{}
	for _, s := range slice {
		newSlice = append(newSlice, f(s))
	}

	return newSlice
}

func Part1(program []int, firstValue, secondValue int) int {
    program[1] = firstValue
    program[2] = secondValue
	insPointer := 0
	for program[insPointer] != INSTRUCTION_HALT {
		if program[insPointer] == INSTRUCTION_ADD {
			posN1 := program[insPointer + 1]
			posN2 := program[insPointer + 2]
			where := program[insPointer + 3]
			program[where] = program[posN1] + program[posN2]

			insPointer += 4
		} else if program[insPointer] == INSTRUCTION_MUL {
			posN1 := program[insPointer + 1]
			posN2 := program[insPointer + 2]
			where := program[insPointer + 3]
			program[where] = program[posN1] * program[posN2]

			insPointer += 4
		}
	}

    return program[0]
}

func Part2(initialProgram, program []int) int {
    program = RestoreProgram(initialProgram)
    target := 19690720
    for noun := 0; noun <= 99; noun++ {
        for verb := 0; verb <= 99; verb++ {
            if Part1(program, noun, verb) == target {
                return 100*noun + verb
            }
            program = RestoreProgram(initialProgram)
        }
    }

    return -1
}

func RestoreProgram(initialProgram []int) []int {
    return slices.Clone[[]int, int](initialProgram)
}
