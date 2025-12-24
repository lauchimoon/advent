package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILEPATH = "./resources/003input.txt"

func main() {
	f, err := os.Open(FILEPATH)
	if err != nil {
		log.Fatalf("failed to open file %s", FILEPATH)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	triangles := []string{}
	for scanner.Scan() {
		triangles = append(triangles, scanner.Text())
	}

	fmt.Println("Part 1:", Part1(triangles))
	fmt.Println("Part 2:", Part2(triangles))
}

func Part1(triangles []string) int {
	total := 0
	for _, tri := range triangles {
		if IsValidTriangle(tri) {
			total++
		}
	}
	return total
}

func Part2(triangles []string) int {
	total := 0
	for i := 0; i < len(triangles); i += 3 {
		three := strings.Fields(strings.Join(triangles[i:i+3], ""))
		for j := 0; j < 3; j++ {
			newTri := strings.Join([]string{three[j], three[j+3], three[j+6]}, " ")
			if IsValidTriangle(newTri) {
				total++
			}
		}
	}
	return total
}

func IsValidTriangle(spec string) bool {
	values := strings.Fields(spec)
	side1, _ := strconv.Atoi(values[0])
	side2, _ := strconv.Atoi(values[1])
	side3, _ := strconv.Atoi(values[2])
	return side1+side2 > side3 && side1+side3 > side2 && side2+side3 > side1
}
