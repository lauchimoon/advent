package main

import (
    "bufio"
    "io"
    "fmt"
    "log"
    "strconv"
    "strings"
    "os"
)

const FILEPATH = "./resources/005input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("couldn't open file '%s'", FILEPATH)
    }
    defer f.Close()

    fileString := LoadFileString(f)
    ranges := LoadRanges(fileString[:185])
    ids := LoadIDs(fileString[186:])

    count := 0
    lenRanges := len(ranges)
    for id := range ids {
        for i := 0; i < lenRanges-1; i += 2 {
            if id >= ranges[i] && id <= ranges[i+1] && !ids[id] {
                ids[id] = true
                count++
            }
        }
    }

    fmt.Println(count)
}

func LoadFileString(f io.Reader) []string {
    s := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        s = append(s, scanner.Text())
    }
    return s
}

func LoadRanges(fileString []string) []uint64 {
    ranges := []uint64{}
    for _, s := range fileString {
        numbersString := strings.Split(s, "-")
        n1, _ := strconv.ParseUint(numbersString[0], 10, 64)
        n2, _ := strconv.ParseUint(numbersString[1], 10, 64)
        ranges = append(ranges, n1, n2)
    }
    return ranges
}

func LoadIDs(fileString []string) map[uint64]bool {
    ids := map[uint64]bool{}
    for _, s := range fileString {
        n, _ := strconv.ParseUint(s, 10, 64)
        ids[n] = false
    }
    return ids
}
