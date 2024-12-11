package day11

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

func (d day) Part1(day int, file string) int {
	stones := load(day, file)

	for i := 0; i < 25; i++ {
		//fmt.Println("Round", i)
		//fmt.Println(stones)
		newStones := make([]string, 0)
		for _, stone := range stones {
			newStones = append(newStones, transform(stone)...)
		}
		stones = newStones
	}

	return len(stones)
}

func load(day int, file string) []string {
	var input []string
	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), func(l string) {
		input = strings.Fields(l)
	})
	return input
}

func transform(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	} else if len(stone)%2 == 0 {
		return getHalfStones(stone)
	} else {
		num := utils.MustAtoi(stone)
		val := num * 2024
		return []string{strconv.Itoa(val)}
	}
}

func getHalfStones(stone string) []string {
	split := len(stone) / 2
	s1 := stone[:split]
	num1 := utils.MustAtoi(s1)
	v1 := strconv.Itoa(num1)
	s2 := stone[split:]
	num2 := utils.MustAtoi(s2)
	v2 := strconv.Itoa(num2)
	return []string{v1, v2}
}
