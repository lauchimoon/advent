package main

import (
    "bufio"
	"fmt"
    "log"
    "os"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	if len(s) == 0 {
		return s, -1
	}

	l := len(s)
	return s[:l-1], s[l-1]
}

const FILEPATH = "./resources/009input.txt"

func main() {
	f, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatalf("failed to open file %s", FILEPATH)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	src := scanner.Text()

	drive := stack{}
	id := 0

	for i, c := range src {
		x := int(c - '0')
		if i%2 == 0 {
			for j := 0; j < x; j++ {
				drive = drive.Push(id)
			}
			id++
		} else {
			for j := 0; j < x; j++ {
				drive = drive.Push(-1)
			}
		}
	}

	blanks := FindBlanks(drive)
	for _, b := range blanks {
		for drive[len(drive)-1] == -1 {
			drive, _ = drive.Pop()
		}

		if len(drive) <= b {
			break
		}

		drive, drive[b] = drive.Pop()
	}

	total := 0
	for i, x := range drive {
		total += i * x
	}
	fmt.Println(total)
}

func FindBlanks(drive stack) stack {
	blanks := stack{}
	for i, x := range drive {
		if x == -1 {
			blanks = blanks.Push(i)
		}
	}
	return blanks
}

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
