package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

const FILEPATH = "./resources/010input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()
    program := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        program = append(program, scanner.Text())
    }

    X := int64(1)
    cycles := int64(0)
    sigStrength := int64(0)
    sumSig := int64(0)

    for _, line := range program {
        if line == "noop" {
            cycles++
            sigStrength = UpdateStrength(sigStrength, cycles, X, &sumSig)
        } else {
            increment, _ := strconv.ParseInt(strings.Fields(line)[1], 10, 64)
            for i := 1; i <= 2; i++ {
                cycles++
                sigStrength = UpdateStrength(sigStrength, cycles, X, &sumSig)
            }
            X += increment
        }
    }
    fmt.Println(sumSig)
}

func UpdateStrength(sigStrength, cycles, X int64, sumSig *int64) int64 {
    if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 ||
       cycles == 180 || cycles == 220 {
           sigStrength = cycles*X
           *sumSig += sigStrength
    }
    return sigStrength
}
