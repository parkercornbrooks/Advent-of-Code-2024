package day10

import (
	"path/filepath"
	"strconv"

	"github.com/parkercornbrooks/advent-of-code-2024/utils"
)

type day struct{}

func New() day {
	return day{}
}

func (d day) Part1(day int, file string) int {
	m := utils.ReadIntoMatrix(filepath.Join("day"+strconv.Itoa(day), file))

	return 0
}
