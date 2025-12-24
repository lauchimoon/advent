package main

import (
	"fmt"
)

type Item struct {
	Cost   int
	Damage int
	Armor  int
}

var (
	weapons = []Item{
		{8, 4, 0}, {10, 5, 0}, {25, 6, 0},
		{40, 7, 0}, {74, 8, 0},
	}
	defense = []Item{
		{0, 0, 0}, {13, 0, 1}, {31, 0, 2},
		{53, 0, 3}, {75, 0, 4}, {102, 0, 5},
	}
	rings = []Item{
		{0, 0, 0},
		{25, 1, 0}, {50, 2, 0}, {100, 3, 0},
		{20, 0, 1}, {40, 0, 2}, {80, 0, 3},
	}
)

const (
	PLAYER_HP = 100
	BOSS_HP = 100
	BOSS_DMG = 8
	BOSS_DEF = 2
)

func main() {
	fmt.Println("Part 1:", Part1())
	fmt.Println("Part 2:", Part2())
}

func Part1() int {
	minCost := 999999
	for _, weapon := range weapons {
		for _, armor := range defense {
			for i, ring1 := range rings {
				for j, ring2 := range rings {
					if i == j && i != 0 {
						continue
					}

					cost := weapon.Cost + ring1.Cost + armor.Cost + ring2.Cost
					playerDmg := weapon.Damage + ring1.Damage + ring2.Damage
					playerDef := armor.Armor + ring1.Armor + ring2.Armor
					if IsPlayerWinner(playerDmg, playerDef) {
						minCost = min(minCost, cost)
					}
				}
			}
		}
	}
	return minCost
}

func IsPlayerWinner(playerDmg, playerDef int) bool {
	playerDamagePerTurn := max(1, playerDmg - BOSS_DEF)
	bossDamagePerTurn := max(1, BOSS_DMG - playerDef)
	playerTurns := BOSS_HP/playerDamagePerTurn
	bossTurns := PLAYER_HP/bossDamagePerTurn
	return playerTurns <= bossTurns
}

func Part2() int {
	maxCost := -999999
	for _, weapon := range weapons {
		for _, armor := range defense {
			for i, ring1 := range rings {
				for j, ring2 := range rings {
					if i == j && i != 0 {
						continue
					}

					cost := weapon.Cost + ring1.Cost + armor.Cost + ring2.Cost
					playerDmg := weapon.Damage + ring1.Damage + ring2.Damage
					playerDef := armor.Armor + ring1.Armor + ring2.Armor
					if !IsPlayerWinner(playerDmg, playerDef) {
						maxCost = max(maxCost, cost)
					}
				}
			}
		}
	}
	return maxCost
}
