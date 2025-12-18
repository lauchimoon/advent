package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

const FILEPATH = "./resources/011input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    paths := LoadPaths(f)
    fmt.Println(CountOut(paths))
}

func LoadPaths(f io.Reader) map[string][]string {
    paths := map[string][]string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, ":")
        origin := split[0]
        dst := strings.Split(strings.TrimLeft(split[1], " "), " ")
        paths[origin] = dst
    }
    return paths
}

func dfs(paths map[string][]string, start string, count *int) int {
    for _, path := range paths[start] {
        if path == "out" {
            (*count)++
        } else {
            dfs(paths, path, count)
        }
    }

    return *count
}

func CountOut(paths map[string][]string) int {
    count := 0
    return dfs(paths, "you", &count)
}
