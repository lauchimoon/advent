package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
)

func main() {
	low := 146810
	high := 612564
	fmt.Printf("Part 1: %d\n", Part1(low, high))
	fmt.Printf("Part 2: %d\n", Part2(low, high))
}

func Part1(low, high int) int {
	total := 0
	for i := low; i <= high; i++ {
		values := slices.Collect[int](maps.Values(GetCount(i)))
		twoAdjacent := slices.Max[[]int, int](values) >= 2
		neverDecreases := NeverDecreases(i)
		if twoAdjacent && neverDecreases {
			total++
		}
	}

	return total
}

func Part2(low, high int) int {
	total := 0
	for i := low; i <= high; i++ {
		values := slices.Collect[int](maps.Values(GetCount(i)))
		twoAdjacent := slices.Contains[[]int, int](values, 2)
		neverDecreases := NeverDecreases(i)
		if twoAdjacent && neverDecreases {
			total++
		}
	}

	return total
}

func AdjacentSame(x int) bool {
	sx := strconv.Itoa(x)
	for i := 0; i < len(sx) - 1; i++ {
		if sx[i] == sx[i + 1] {
			return true
		}
	}
	return false
}

func GetCount(x int) map[int]int {
	sx := strconv.Itoa(x)
	count := map[int]int{}
	for _, c := range sx {
		count[int(c - '0')]++
	}

	return count
}

func NeverDecreases(x int) bool {
	sx := strconv.Itoa(x)
	for i := 0; i < len(sx) - 1; i++ {
		if sx[i] > sx[i + 1] {
			return false
		}
	}
	return true
}
