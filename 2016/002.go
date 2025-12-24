package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const FILEPATH = "./resources/002input.txt"

var (
	keypad1 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	keypad2 = [][]int{
		{-1, -1, '1', -1, -1},
		{-1, '2', '3', '4', -1},
		{'5', '6', '7', '8', '9'},
		{-1, 'A', 'B', 'C', -1},
		{-1, -1, 'D', -1, -1},
	}
)

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

	scanner := bufio.NewScanner(f)
	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	fmt.Println("Part 1:", Part1(instructions))
	fmt.Println("Part 2:", Part2(instructions))
}

func Part1(instructions []string) string {
	row, col := 1, 1
	code := ""
	for _, ins := range instructions {
		dig, newRow, newCol := ParseInstruction1(ins, row, col)
		code += string(dig + '0')
		row, col = newRow, newCol
	}
	return code
}

func ParseInstruction1(instruction string, prevRow, prevCol int) (int, int, int) {
	row, col := prevRow, prevCol
	for _, c := range instruction {
		switch c {
		case 'U':
			if row > 0 {
				row--
			}
			break
		case 'R':
			if col < len(keypad1[0]) - 1 {
				col++
			}
			break
		case 'D':
			if row < len(keypad1) - 1 {
				row++
			}
			break
		case 'L':
			if col > 0 {
				col--
			}
			break
		}
	}

	return keypad1[row][col], row, col
}

func Part2(instructions []string) string {
	row, col := 3, 0
	code := ""
	for _, ins := range instructions {
		char, newRow, newCol := ParseInstruction2(ins, row, col)
		code += string(char)
		row, col = newRow, newCol
	}
	return code
}

func ParseInstruction2(instruction string, prevRow, prevCol int) (int, int, int) {
	row, col := prevRow, prevCol
	for _, c := range instruction {
		switch c {
		case 'U':
			if row > 0 && !IsEdge(row - 1, col) {
				row--
			}
			break
		case 'R':
			if col < len(keypad2) - 1 && !IsEdge(row, col + 1) {
				col++
			}
			break
		case 'D':
			if row < len(keypad2[0]) - 1 && !IsEdge(row + 1, col) {
				row++
			}
			break
		case 'L':
			if col > 0 && !IsEdge(row, col - 1) {
				col--
			}
			break
		}
	}
	return keypad2[row][col], row, col
}

func IsEdge(row, col int) bool {
	return keypad2[row][col] == -1
}
