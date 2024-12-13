package day13

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
	linefn := func(line string) {}

	utils.ReadInput(filepath.Join("day"+strconv.Itoa(day), file), linefn)

	return 0
}
