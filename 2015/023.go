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
	Op  string
	Reg string
	Arg int
}

var (
	registers = map[string]int{"a": 0, "b": 0}
)

const FILEPATH = "./resources/023input.txt"

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
	program := LoadProgram(lines)
	Interpret(program)
	return registers["b"]
}

func Part2(lines []string) int {
	registers["a"] = 1
	registers["b"] = 0
	return Part1(lines)
}

func LoadProgram(lines []string) []Instruction {
	program := []Instruction{}
	for _, line := range lines {
		ins := Instruction{}
		fields := strings.Fields(line)
		op := fields[0]
		reg := fields[1]
		ins.Op = op
		ins.Reg = reg

		// No jump
		ins.Arg = -1

		if op == "jmp" {
			ins.Reg = "?"
			ins.Arg, _ = strconv.Atoi(fields[1])
		}
		if op == "jie" || op == "jio" {
			ins.Reg = strings.Trim(reg, ",")
			ins.Arg, _ = strconv.Atoi(fields[2])
		}

		program = append(program, ins)
	}

	return program
}

func Interpret(program []Instruction) {
	for pc := 0; pc < len(program); pc++ {
		ins := program[pc]
		switch ins.Op {
		case "hlf":
			registers[ins.Reg] /= 2
			break
		case "tpl":
			registers[ins.Reg] *= 3
			break
		case "inc":
			registers[ins.Reg]++
			break
		case "jmp":
			pc += ins.Arg - 1
			break
		case "jie":
			if registers[ins.Reg] % 2 == 0 {
				pc += ins.Arg - 1
			}
			break
		case "jio":
			if registers[ins.Reg] == 1 {
				pc += ins.Arg - 1
			}
			break
		}
	}
}
