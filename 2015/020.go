package main

import (
	"fmt"
	"math"
)

const TARGET = 33100000

func main() {
	fmt.Printf("Part 1: %d\n", Part1())
	fmt.Printf("Part 2: %d\n", Part2())
}

func Part1() int {
	house := 1
	for SumOfPresents1(house) < TARGET {
		house++
	}
	return house
}

func Part2() int {
	house := 1
	for SumOfPresents2(house) < TARGET {
		house++
	}
	return house
}

func SumOfPresents1(n int) int {
	sum := 0
	top := int(math.Sqrt(float64(n))) + 1
	for i := 1; i <= top; i++ {
		if n % i == 0 {
			sum += i
			sum += n/i
		}
	}

	return sum*10
}

func SumOfPresents2(n int) int {
	sum := 0
	top := int(math.Sqrt(float64(n))) + 1
	for i := 1; i <= top; i++ {
		if n % i == 0 {
			if (i <= 50) {
				sum += n/i
			}

			if (n/i <= 50) {
				sum += i
			}
		}
	}

	return sum*11
}
