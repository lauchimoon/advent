package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

const FILEPATH = "./resources/004input.txt"

func main() {
    f, err := os.Open(FILEPATH)
    if err != nil {
        log.Fatalf("failed to open file %s", FILEPATH)
    }
    defer f.Close()

    passports := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        passports = append(passports, scanner.Text())
    }

    fmt.Printf("Part 1: %d\n", Part1(passports))
}

func Part1(passports []string) int {
    total := 0
    passportData := ""
    for _, passport := range passports {
        passportData += passport + " "
        if passport == "" {
            if IsValidPassport(passportData) {
                total++
            }
            passportData = ""
        }
    }

    // Check the last one too
    if IsValidPassport(passportData) {
        total++
    }

    return total
}

func IsValidPassport(data string) bool {
    // cid is optional
    requiredFields := []string{
        "byr", "iyr", "eyr", "hgt",
        "hcl", "ecl", "pid",
    }

    fields := strings.Fields(data)
    availFields := map[string]bool{}
    for _, fieldData := range fields {
        field := strings.Split(fieldData, ":")[0]
        availFields[field] = true
    }

    for _, field := range requiredFields {
        _, ok := availFields[field]
        if !ok {
            return false
        }
    }

    return true
}
