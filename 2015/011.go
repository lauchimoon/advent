package main

import (
	"fmt"
)

func main() {
	input := "vzbxkghb"
	fmt.Printf("Part 1: %s\n", Part1(input))
	input = Part1(input)
	fmt.Printf("Part 2: %s\n", Part1(IncrementString(input)))
}

func Part1(input string) string {
	for !IsValid(input) {
		input = IncrementString(input)
	}
	return input
}

func IncrementString(s string) string {
	chars := []byte(s)
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] == 'z' {
			chars[i] = 'a'
		} else {
			chars[i]++
			return string(chars)
		}
	}

	return "a" + string(chars)
}

func IsValid(pass string) bool {
	return IncreasingStraight(pass) &&
		   !Confusing(pass) &&
		   LetterPairs(pass)
}

func IncreasingStraight(pass string) bool {
	for i := 0; i < len(pass) - 2; i++ {
		if pass[i+1] - pass[i] == 1 &&
		   pass[i+2] - pass[i+1] == 1 {
			   return true
	    }
	}

	return false
}

func Confusing(pass string) bool {
	for _, c := range pass {
		if c == 'i' || c == 'l' || c == 'o' {
			return true
		}
	}

	return false
}

func LetterPairs(pass string) bool {
	var firstFound byte = ' '
	pairCount := 0

	for i := 0; i < len(pass) - 1; i++ {
		if pass[i] == pass[i+1] {
			pairCount++
			if pass[i] != firstFound {
				firstFound = pass[i]
			} else {
				return false
			}
		}
	}

	return pairCount >= 2
}
