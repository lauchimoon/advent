package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	N_STEPS = 100

	FILEPATH = "./resources/018input.txt"
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

	grid := LoadGrid(lines)
	fmt.Printf("Part 1: %d\n", Part1(grid))
	grid = LoadGrid(lines)
	fmt.Printf("Part 2: %d\n", Part2(grid))
}

func LoadGrid(lines []string) [][]int {
	height := len(lines)
	width := len(lines[0])

	grid := make([][]int, height)
	for i := 0; i < width; i++ {
		grid[i] = make([]int, width)
	}

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '#' {
				grid[y][x] = 1
			} else {
				grid[y][x] = 0
			}
		}
	}
	return grid
}

func Part1(grid [][]int) int {
	step := 0
	for step < N_STEPS {
		grid = NextState1(grid)
		step++
	}
	return CountOn(grid)
}

func Part2(grid [][]int) int {
	step := 0
	for step < N_STEPS {
		grid = NextState2(grid)
		step++
	}
	return CountOn(grid)
}

func NextState1(state [][]int) [][]int {
	newState := make([][]int, len(state))
	for i := 0; i < len(state); i++ {
		newState[i] = make([]int, len(state))
	}

	for y := range state {
		for x := range state[y] {
			onCount := CountNeighbours(state, y, x, 1)
			switch state[y][x] {
			case 0:
				if onCount == 3 {
					newState[y][x] = 1
				}
				break
			case 1:
				if onCount == 2 || onCount == 3 {
					newState[y][x] = 1
				}
				break
			}
		}
	}

	return newState
}

func NextState2(state [][]int) [][]int {
	height := len(state)
	width := len(state[0])

	newState := make([][]int, height)
	for i := 0; i < len(state); i++ {
		newState[i] = make([]int, width)
	}

	newState[0][0] = 1
	newState[0][width - 1] = 1
	newState[height - 1][0] = 1
	newState[height - 1][width - 1] = 1

	for y := range state {
		for x := range state[y] {
			// Ignore corners
			if (y == 0 && x == 0) || (y == height - 1 && x == 0) || 
				(y == 0 && x == width - 1) || (y == height - 1 && x == width - 1) {
					continue
			}

			onCount := CountNeighbours(state, y, x, 1)
			switch state[y][x] {
			case 0:
				if onCount == 3 {
					newState[y][x] = 1
				}
				break
			case 1:
				if onCount == 2 || onCount == 3 {
					newState[y][x] = 1
				}
				break
			}
		}
	}

	return newState
}

func CountNeighbours(state [][]int, y, x, cellState int) int {
	neighbours := []int{}
	directions := []int{-1, 0, 1}

	for _, dy := range directions {
		for _, dx := range directions {
			if dy == 0 && dx == 0 {
				continue
			}

			newY, newX := y + dy, x + dx
			if newY >= 0 && newY < len(state) &&
				newX >= 0 && newX < len(state) &&
				state[newY][newX] == cellState {
					neighbours = append(neighbours, state[newY][newX])
			}
		}
	}

	return len(neighbours)
}

func CountOn(state [][]int) int {
	total := 0
	for y := range state {
		for x := range state[y] {
			if state[y][x] == 1 {
				total++
			}
		}
	}

	return total
}
