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

	return run(stones, 25)
}

func load(day int, file string) []string {
	var input []string
	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), func(l string) {
		input = strings.Fields(l)
	})
	return input
}

func getHalfStones(stone string) (string, string) {
	split := len(stone) / 2
	s1 := stone[:split]
	num1 := utils.MustAtoi(s1)
	v1 := strconv.Itoa(num1)
	s2 := stone[split:]
	num2 := utils.MustAtoi(s2)
	v2 := strconv.Itoa(num2)
	return v1, v2
}
