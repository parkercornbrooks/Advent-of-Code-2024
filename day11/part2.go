package day11

import (
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

func (d day) Part2(day int, file string) int {
	stones := load(day, file)
	return run(stones, 75)
}

func run(stones []string, blinks int) int {
	stoneMap := map[string]int{}
	for _, stone := range stones {
		stoneMap[stone] += 1
	}

	for i := 0; i < blinks; i++ {
		newStones := map[string]int{}
		for stone, count := range stoneMap {
			if stone == "0" {
				newStones["1"] += count
			} else if len(stone)%2 == 0 {
				a, b := getHalfStones(stone)
				newStones[a] += count
				newStones[b] += count
			} else {
				num := utils.MustAtoi(stone)
				val := num * 2024
				newStones[strconv.Itoa(val)] += count
			}
		}
		stoneMap = newStones
	}

	total := 0
	for _, v := range stoneMap {
		total += v
	}

	return total
}
