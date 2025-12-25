package main

import (
	"cmp"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
	"slices"
	"unicode"
)

type Pair struct {
	First  rune
	Second int
}

const FILEPATH = "./resources/004input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

	rooms := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rooms = append(rooms, scanner.Text())
	}

	part1Result, validRooms := Part1(rooms)
	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", Part2(validRooms))
}

func Part1(rooms []string) (int, []string) {
	total := 0
	validRooms := []string{}
	for _, room := range rooms {
		check, id := IsValidRoom(room)
		if check {
			validRooms = append(validRooms, room)
			total += id
		}
	}
	return total, validRooms
}

func Part2(validRooms []string) int {
	for _, room := range validRooms {
		fields := strings.Split(room[:strings.Index(room, "[")], "-")
		encryptedName := strings.Join(fields[:len(fields)-1], "-")
		shift, _ := strconv.Atoi(fields[len(fields)-1])
		decrypted := Decrypt(encryptedName, shift)

		if decrypted == "northpole object storage" {
			return shift
		}
	}

	return -1
}

func Decrypt(s string, shift int) string {
	r := []rune(s)
	for i := range s {
		if r[i] == '-' {
			r[i] = ' '
		} else if unicode.IsLetter(r[i]) {
			r[i] = rune(int(r[i] - 'a') + shift) % 26 + 'a'
		}
	}

	return string(r)
}

func IsValidRoom(room string) (bool, int) {
	openBracketIdx := strings.Index(room, "[") + 1
	checksum := room[openBracketIdx:openBracketIdx + 5]

	info := strings.Split(room, "-")
	id := stripId(info[len(info) - 1])

	freqLetters := map[rune]int{}
	for _, s := range info[:len(info) - 1] {
		for _, c := range s {
			freqLetters[c]++
		}
	}

	fiveFrequent := []Pair{}
	for k, v := range freqLetters {
		fiveFrequent = append(fiveFrequent, Pair{k, v})
	}

	slices.SortFunc(fiveFrequent, func(a, b Pair) int {
		if v := cmp.Compare(b.Second, a.Second); v != 0 {
			return v
		}
		return cmp.Compare(a.First, b.First)
	})

	fiveFrequent = fiveFrequent[:5]
	return ProduceChecksum(fiveFrequent) == SortString(checksum), id
}

func stripId(s string) int {
	id, _ := strconv.Atoi(s[:strings.Index(s, "[")])
	return id
}

func ProduceChecksum(frequent []Pair) string {
	checksum := make([]rune, 5)
	for i, p := range frequent {
		checksum[i] = p.First
	}

	return SortString(string(checksum))
}

func SortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
