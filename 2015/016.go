package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILEPATH = "./resources/016input.txt"

var (
	targetSue = map[string]int{
		"children": 3, "cats": 7, "samoyeds": 2,
		"pomeranians": 3, "akitas": 0, "vizslas": 0,
		"goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1,
	}
)

func main() {
	f, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatalf("failed to open file %s", FILEPATH)
	}
	defer f.Close()

	sues := []map[string]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sues = append(sues, ParseLine(scanner.Text()))
	}

	fmt.Printf("Part 1: %d\n", Part1(sues))
	fmt.Printf("Part 2: %d\n", Part2(sues))
}

func Part1(sues []map[string]int) int {
	for i, sue := range sues {
		allMatch := true
		for k, v := range sue {
			if targetSue[k] != v {
				allMatch = false
				break
			}
		}

		if allMatch {
			return i + 1
		}
	}
	return -1
}

func Part2(sues []map[string]int) int {
	for i, sue := range sues {
		allMatch := true

		for _, check := range []string{"cats", "trees"} {
			if count, ok := sue[check]; ok {
				if count <= targetSue[check] {
					allMatch = false
				}
				delete(sue, check)
			}
		}

		for _, check := range []string{"pomeranians", "goldfish"} {
			if count, ok := sue[check]; ok {
				if count >= targetSue[check] {
					allMatch = false
				}
				delete(sue, check)
			}
		}

		for k, v := range sue {
			if targetSue[k] != v {
				allMatch = false
				break
			}
		}

		if allMatch {
			return i + 1
		}
	}

	return -1
}

func ParseLine(s string) map[string]int {
	m := map[string]int{}
	var category1, category2, category3 string
	var n1, n2, n3, dummyN int
	fmt.Sscanf(s, "Sue %d: %s %d, %s %d, %s %d", &dummyN, &category1, &n1,
											 	 &category2, &n2, &category3, &n3)

	category1 = strings.Trim(category1, ":")
	category2 = strings.Trim(category2, ":")
	category3 = strings.Trim(category3, ":")
	m[category1] = n1
	m[category2] = n2
	m[category3] = n3
	return m
}
