package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	key := "yzbqklnj"
	fmt.Printf("Part 1: %d\n", Part1(key))
	fmt.Printf("Part 2: %d\n", Part2(key))
}

func Part1(key string) int {
	i := 1

	for true {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("%s%d", key, i))
		sh := fmt.Sprintf("%x", h.Sum(nil))
		firstFive := sh[0:5]
		if firstFive == "00000" {
			break
		}

		i++
	}

	return i
}

func Part2(key string) int {
	i := 1

	for true {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("%s%d", key, i))
		sh := fmt.Sprintf("%x", h.Sum(nil))
		firstSix := sh[0:6]
		if firstSix == "000000" {
			break
		}

		i++
	}

	return i
}
