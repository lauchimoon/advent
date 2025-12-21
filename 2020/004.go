package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "unicode"
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

    fmt.Printf("Part 1: %d\n", Solve(passports, 1))
    fmt.Printf("Part 2: %d\n", Solve(passports, 2))
}

func Solve(passports []string, part int) int {
    total := 0
    passportData := ""
    for _, passport := range passports {
        passportData += passport + " "
        if passport == "" {
            if IsValidPassport(passportData, part) {
                total++
            }
            passportData = ""
        }
    }

    // Check the last one too
    if IsValidPassport(passportData, part) {
        total++
    }

    return total
}

func IsValidPassport(data string, part int) bool {
    // cid is optional
    requiredFields := []string{
        "byr", "iyr", "eyr", "hgt",
        "hcl", "ecl", "pid",
    }

    fields := strings.Fields(data)
    availFields := map[string]bool{}
    for _, fieldData := range fields {
        if part == 1 {
            field := strings.Split(fieldData, ":")[0]
            availFields[field] = true
        } else {
            info := strings.Split(fieldData, ":")
            field := info[0]
            data := info[1]
            if IsValidField(field, data) {
                availFields[field] = true
            } else {
                return false
            }
        }
    }

    for _, field := range requiredFields {
        _, ok := availFields[field]
        if !ok {
            return false
        }
    }

    return true
}

func IsValidField(field, data string) bool {
    switch field {
    case "byr":
        dataInt, _ := strconv.Atoi(data)
        return dataInt >= 1920 && dataInt <= 2002
    case "iyr":
        dataInt, _ := strconv.Atoi(data)
        return dataInt >= 2010 && dataInt <= 2020
    case "eyr":
        dataInt, _ := strconv.Atoi(data)
        return dataInt >= 2020 && dataInt <= 2030
    case "hgt":
        dataInt, _ := strconv.Atoi(data[:len(data)-2]) // remove cm or in
        if strings.Contains(data, "cm") {
            return dataInt >= 150 && dataInt <= 193
        } else if strings.Contains(data, "in") {
            return dataInt >= 59 && dataInt <= 76
        } else {
            return false
        }
    case "hcl":
        return validHairColor(data)
    case "ecl":
        return validEyeColor(data)
    case "pid":
        return validPassportId(data)
    }
    return true
}

func validHairColor(data string) bool {
    if !strings.Contains(data, "#") {
        return false
    }

    noHash := data[1:]
    if len(noHash) != 6 {
        return false
    }

    for _, c := range noHash {
        if !unicode.IsDigit(c) && (unicode.IsLetter(c) && (c < 'a' || c > 'f')) {
            return false
        }
    }

    return true
}

func validEyeColor(data string) bool {
    possibleColors := strings.Fields("amb blu brn gry grn hzl oth")
    for _, color := range possibleColors {
        if data == color {
            return true
        }
    }

    return false
}

func validPassportId(data string) bool {
    if len(data) != 9 {
        return false
    }

    for _, c := range data {
        if !unicode.IsDigit(c) {
            return false
        }
    }

    return true
}
