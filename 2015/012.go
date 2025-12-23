package main

import (
    "bufio"
	"encoding/json"
    "fmt"
    "log"
    "os"
	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`-?[0-9]+`)
)

const FILEPATH = "./resources/012input.txt"

func main() {
	f, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatalf("failed to open file %s", FILEPATH)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	document := scanner.Text()

	fmt.Printf("Part 1: %d\n", GetNumberSum(document))
	fmt.Printf("Part 2: %d\n", GetNonRedSum(document))
}

func GetNumberSum(s string) int {
	total := 0
	for _, num := range re.FindAllString(s, -1) {
		x, _ := strconv.Atoi(num)
		total += x
	}

	return total
}

func GetNonRedSum(s string) int {
	if !regexp.MustCompile("red").MatchString(s) {
		return GetNumberSum(s)
	}

	var obj map[string]interface{}
	err := json.Unmarshal([]byte(s), &obj)

	// obj is Array ([])
	if err != nil {
		var arr []interface{}
		json.Unmarshal([]byte(s), &arr)

		var arrTotal int
		for _, v := range arr {
			s, _ := json.Marshal(v)
			arrTotal += GetNonRedSum(string(s))
		}

		return arrTotal
	}

	// obj is Object ({})
	for _, v := range obj {
		s, ok := v.(string)
		if ok && s == "red" {
			return 0
		}
	}

	var total int
	for _, v := range obj {
		s, _ := json.Marshal(v)
		total += GetNonRedSum(string(s))
	}
	return total
}
