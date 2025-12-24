package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILEPATH = "./resources/001input.txt"

type Instruction struct {
	Direction byte
	Move      int
}

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	instructions := LoadInstructions(scanner.Text())
	fmt.Println("Part 1:", Part1(instructions))
	fmt.Println("Part 2:", Part2(instructions))
}

func LoadInstructions(s string) []Instruction {
	instructions := []Instruction{}
	for _, ins := range strings.Split(s, ", ") {
		dir := ins[0]
		mov, _ := strconv.Atoi(ins[1:])
		instructions = append(instructions, Instruction{dir, mov})
	}
	return instructions
}

func Part1(instructions []Instruction) int {
	// row col format
	dirs := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	row, col := 0, 0
	d := 0
	for _, ins := range instructions {
		if ins.Direction == 'R' {
			d = (d + 1) % 4
		} else {
			d = (d + 3) % 4
		}
		row += ins.Move*dirs[d][0]
		col += ins.Move*dirs[d][1]
	}

	return Manhattan(0, 0, row, col)
}

func Part2(instructions []Instruction) int {
	// row col format
	dirs := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	visited := map[[2]int]bool{
		{0, 0}: true,
	}
	row, col := 0, 0
	d := 0
	for _, ins := range instructions {
		if ins.Direction == 'R' {
			d = (d + 1) % 4
		} else {
			d = (d + 3) % 4
		}

		for i := 0; i < ins.Move; i++ {
			row += dirs[d][0]
			col += dirs[d][1]
			if visited[[2]int{row, col}] {
				return Manhattan(0, 0, row, col)
			}
			visited[[2]int{row, col}] = true
		}
	}

	return -1
}

func Manhattan(x1, x2, y1, y2 int) int {
	return Abs(x1 - x2) + Abs(y1 - y2)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
