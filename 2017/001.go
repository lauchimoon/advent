package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const FILEPATH = "./resources/001input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	seq := scanner.Text()

	fmt.Println("Part 1:", Captcha(seq, 1))
	fmt.Println("Part 2:", Captcha(seq, len(seq)/2))
}

func Captcha(seq string, shift int) int {
	total := 0
	lenSequence := len(seq)

	for i := 0; i < lenSequence; i++ {
		if seq[i] == seq[(i + shift) % lenSequence] {
			total += int(seq[i] - '0')
		}
	}

	return total
}
