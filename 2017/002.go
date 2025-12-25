package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const FILEPATH = "./resources/002input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

	scanner := bufio.NewScanner(f)
	rows := []string{}
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	fmt.Println("Part 1:", Part1(rows))
	fmt.Println("Part 2:", Part2(rows))
}

func Part1(rows []string) int {
	total := 0
	for _, row := range rows {
		nums := []int{}
		for _, s := range strings.Fields(row) {
			x, _ := strconv.Atoi(s)
			nums = append(nums, x)
		}
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})

		total += nums[len(nums) - 1] - nums[0]
	}

	return total
}

func Part2(rows []string) int {
	total := 0
	for _, row := range rows {
		nums := []int{}
		for _, s := range strings.Fields(row) {
			x, _ := strconv.Atoi(s)
			nums = append(nums, x)
		}

		for _, i := range nums {
			for _, j := range nums {
				if i == j {
					continue
				}

				if j != 0 && i % j == 0 {
					total += i/j
					break
				}
			}
		}
	}

	return total
}
