package main

import (
	"fmt"
)

const (
	TOP = 100
	N_PROPERTIES = 4
)

func main() {
	ingredients := [][]int64{
		{3, 0, 0, -3, 2},	// Sugar
		{-3, 3, 0, 0, 9},	// Sprinkles
		{-1, 0, 4, 0, 1},	// Candy
		{0, 0, -2, 2, 8},	// Chocolate
	}

	fmt.Printf("Part 1: %d\n", Part1(ingredients))
	fmt.Printf("Part 2: %d\n", Part2(ingredients))
}

func Part1(ingredients [][]int64) int64 {
	var score int64 = 0
	for sugar := 0; sugar <= TOP; sugar++ {
		for sprinkles := 0; sprinkles <= TOP-sugar; sprinkles++ {
			for candy := 0; candy <= TOP - sugar - sprinkles; candy++ {
				chocolate := TOP - sugar - sprinkles - candy
				newScore, _ := CalculateScore(ingredients, []int64{
					int64(sugar), int64(sprinkles), int64(candy), int64(chocolate),
				})
				score = max(newScore, score)
			}
		}
	}

	return score
}

func Part2(ingredients [][]int64) int64 {
	var score int64 = 0
	for sugar := 0; sugar <= TOP; sugar++ {
		for sprinkles := 0; sprinkles <= TOP-sugar; sprinkles++ {
			for candy := 0; candy <= TOP - sugar - sprinkles; candy++ {
				chocolate := TOP - sugar - sprinkles - candy
				newScore, calories := CalculateScore(ingredients, []int64{
					int64(sugar), int64(sprinkles), int64(candy), int64(chocolate),
				})

				if calories == 500 {
					score = max(newScore, score)
				}
			}
		}
	}

	return score
}

func CalculateScore(ingredients [][]int64, amounts []int64) (int64, int64) {
	var capacity, durability, flavor, texture, calories int64 = 0, 0, 0, 0, 0
	for property := 0; property < N_PROPERTIES; property++ {
		capacity += amounts[property]*ingredients[property][0]
		durability += amounts[property]*ingredients[property][1]
		flavor += amounts[property]*ingredients[property][2]
		texture += amounts[property]*ingredients[property][3]
		calories += amounts[property]*ingredients[property][4]
	}

	score := capacity*durability*flavor*texture
	if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
		return 0, 0
	}

	return score, calories
}
