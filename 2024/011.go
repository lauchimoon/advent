package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

const FILEPATH = "./resources/011input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f); scanner.Scan()
    src := scanner.Text()
    newSrc := src
    ss := strings.Fields(newSrc)
    newSrcSlice := []string{}

    for i := 0; i < 25; i++ {
        newSrcSlice = []string{}
        for _, s := range ss {
            if s == "0" {
                newSrcSlice = append(newSrcSlice, "1")
            } else if len(s) % 2 == 0 {
                mid := len(s)/2
                x, _ := strconv.ParseInt(s[:mid], 10, 64)
                newSrcSlice = append(newSrcSlice, strconv.FormatInt(x, 10))
                x, _ = strconv.ParseInt(s[mid:], 10, 64)
                newSrcSlice = append(newSrcSlice, strconv.FormatInt(x, 10))
            } else {
                x, _ := strconv.ParseInt(s, 10, 64)
                newSrcSlice = append(newSrcSlice, strconv.FormatInt(x*2024, 10))
            }
        }
        newSrc = strings.Join(newSrcSlice, " ")
        ss = strings.Fields(newSrc)
    }
    fmt.Println(len(newSrcSlice))
}
