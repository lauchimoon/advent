package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

const FILEPATH = "./resources/008input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Scan()
    instruction := InstructionToSlice(scanner.Text())

    // Omit whitespace
    scanner.Scan()

    network := BuildNetwork(scanner)
    current := "AAA"
    i := 0
    steps := 0
    lenInstruction := len(instruction)
    for current != "ZZZ" {
        current = network[current][instruction[i]]
        i = (i + 1) % lenInstruction
        steps++
    }
    fmt.Println(steps)
}

func InstructionToSlice(inst string) []int {
    slice := []int{}
    for _, c := range inst {
        if c == 'L' {
            slice = append(slice, 0)
        } else {
            slice = append(slice, 1)
        }
    }
    return slice
}

func BuildNetwork(scanner *bufio.Scanner) map[string][]string {
    network := map[string][]string{}
    for scanner.Scan() {
        split := strings.Split(scanner.Text(), " = ")
        node := split[0]
        nextNodes := split[1]

        network[node] = TupleToSlice(nextNodes[1:len(nextNodes)-1])
    }
    return network
}

func TupleToSlice(tuple string) []string {
    slice := make([]string, 2)
    split := strings.Split(tuple, ", ")
    slice[0] = split[0]
    slice[1] = split[1]

    return slice
}
