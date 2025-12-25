package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILEPATH = "./resources/008input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for	scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	registers := LoadRegisters(lines)
	fmt.Println("Part 1:", Part1(registers, lines))

	// Reset register values
	registers = LoadRegisters(lines)
	fmt.Println("Part 2:", Part2(registers, lines))
}

func LoadRegisters(lines []string) map[string]int {
	registers := map[string]int{}
	for _, line := range lines {
		registers[strings.Fields(line)[0]] = 0
	}
	return registers
}

func Part1(registers map[string]int, lines []string) int {
	largest := -99999999
	for _, line := range lines {
		fields := strings.Fields(line)
		reg := fields[0]
		op := fields[1]
		opVal, _ := strconv.Atoi(fields[2])
		cond := strings.Join(fields[4:], " ")
		registers[reg] = ParseInstruction(registers, reg, op, opVal, cond)
	}

	for _, v := range registers {
		largest = max(largest, v)
	}
	return largest
}

func Part2(registers map[string]int, lines []string) int {
	largest := -99999999
	for _, line := range lines {
		fields := strings.Fields(line)
		reg := fields[0]
		op := fields[1]
		opVal, _ := strconv.Atoi(fields[2])
		cond := strings.Join(fields[4:], " ")
		registers[reg] = ParseInstruction(registers, reg, op, opVal, cond)
		largest = max(largest, registers[reg])
	}

	return largest
}

func ParseInstruction(registers map[string]int, reg, op string, opVal int, cond string) int {
	condFields := strings.Fields(cond)
	condReg := condFields[0]
	condOp := condFields[1]
	condVal, _ := strconv.Atoi(condFields[2])

	switch condOp {
	case "==":
		if registers[condReg] == condVal {
			return ParseOp(registers, reg, op, opVal)
		}
		break
	case "!=":
		if registers[condReg] != condVal {
			return ParseOp(registers, reg, op, opVal)
		}
		break
	case ">=":
		if registers[condReg] >= condVal {
			return ParseOp(registers, reg, op, opVal)
		}
		break
	case "<=":
		if registers[condReg] <= condVal {
			return ParseOp(registers, reg, op, opVal)
		}
		break
	case ">":
		if registers[condReg] > condVal {
			return ParseOp(registers, reg, op, opVal)
		}
		break
	case "<":
		if registers[condReg] < condVal {
			return ParseOp(registers, reg, op, opVal)
		}
		break
	}

	return registers[reg]
}

func ParseOp(registers map[string]int, reg, op string, opVal int) int {
	if op == "inc" {
		return registers[reg] + opVal
	} else {
		return registers[reg] - opVal
	}
}
