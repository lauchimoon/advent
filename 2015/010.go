package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "3113322113"
	fmt.Printf("Part 1: %d\n", Calculate(input, 40))
	fmt.Printf("Part 2: %d\n", Calculate(input, 50))
}

func Calculate(input string, top int) int {
	for i := 0; i < top; i++ {
		input = NextState(input)
	}
	return len(input)
}

func NextState(s string) string {
	var nextS strings.Builder
	nextS.Grow(2*len(s))

	for i := 0; i < len(s); i++ {
		count := 1
		for i + 1 < len(s) && s[i] == s[i + 1] {
			count++
			i++
		}

		nextS.WriteByte(byte(count + '0'))
		nextS.WriteByte(s[i])
	}

	return nextS.String()
}
